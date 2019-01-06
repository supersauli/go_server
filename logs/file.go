package logs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

type LogFile struct {
	sync.RWMutex

	// file name
	FileName string
	// 文件后缀 .log
	Suffix string
	//
	fileIndex int
	// 文件路径名
	filePathName string
	// maxlinelimit
	MaxLines int
	curLines int

	MaxFileNum int
	CurFileNum int

	// file size limt
	MaxFileSize int
	curFileSize int

	//
	FilePath string
	//data formt
	// 最大天数
	MaxDays int
	// 等级
	Level int
	// 当前文件
	fileWriteNow *os.File
	// 文件权限
	Perm string
	// 日志文件名格式
	FileDatePattern string

	// 刷新方式
	// 每年刷新
	isYearRefresh bool
	// 每月刷新
	isMonthRefresh bool
	// 每日刷新
	isDailyRefresh bool
	// 每小时刷新
	isHourRefresh bool

	// 文件日期
	fileDate time.Time
	//
	fileYear  int
	fileMonth int
	fileDay   int
	fileHour  int
}

func createLog() Logger {
	f := &LogFile{
		FileDatePattern: "-2006-2-1-15-",
		Perm:            "0666",
		Level:           LogLevelTrace,
		Suffix:          ".log",
		MaxFileNum:      9999,
		MaxFileSize:     10000000,
		MaxLines:        10000000,
		FilePath:        "./logs/",
	}
	return f
}

func (f *LogFile) Log(level int, format string, data ...interface{}) {
	if level < f.Level {
		return
	}

	timeNow := time.Now()
	y, m, d := timeNow.Date()
	h, _, _ := timeNow.Clock()
	if f.canSwitchFile(y, int(m), d, h) {
		f.switchNewFile(timeNow)
	}

	logMsg := GetLogHeard(timeNow, level) + fmt.Sprintf(format, data...)
	writeSize, err := f.fileWriteNow.Write([]byte(logMsg))
	if err == nil {
		f.curFileSize += writeSize
		f.curLines += 1
	}
	//f.Unlock()
}

func (f *LogFile) Stop() {
	f.fileWriteNow.Close()
}

func (f *LogFile) Init(jsonConfig string) error {
	err := json.Unmarshal([]byte(jsonConfig), f)
	if err != nil {
		return err
	}
	if len(f.FileDatePattern) == 0 {
		return errors.New("log name not set")
	}
	f.initTimeRef()
	return f.createFile()
}

func (f *LogFile) createFile() error {
	file, err := f.createLogFile()
	if err != nil {
		return err
	}

	if f.fileWriteNow != nil {
		f.fileWriteNow.Close()
	}
	f.fileWriteNow = file
	return f.initFile()
}

func (f *LogFile) createLogFile() (*os.File, error) {
	perm, err := strconv.ParseInt(f.Perm, 8, 64)
	if err != nil {
		return nil, err
	}
	mode := os.FileMode(perm)
	//f.fileName = f.getFileName(time.Now())
	filePath := path.Dir(f.FilePath)
	os.MkdirAll(filePath, mode)
	f.filePathName = f.FilePath + f.FileName + f.Suffix
	fd, err := os.OpenFile(f.filePathName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, mode)
	if err != nil {
		return nil, err
	}
	return fd, err
}

func (f *LogFile) initFile() error {
	fd := f.fileWriteNow
	fileInfo, err := fd.Stat()
	if err != nil {
		return err
	}

	f.curFileSize = int(fileInfo.Size())
	var timeValue time.Time
	var timeNow = time.Now()
	if f.curFileSize == 0 {
		timeValue = timeNow
	} else {
		timeValue = fileInfo.ModTime()
	}

	f.fileDate = timeValue
	f.fileYear = timeValue.Year()
	f.fileMonth = int(timeValue.Month())
	f.fileDay = timeValue.Day()
	f.fileHour = timeValue.Hour()
	f.curLines = f.getLines()
	if f.canSwitchFile(timeNow.Year(), int(timeNow.Month()), timeNow.Day(), timeNow.Hour()) {
		f.switchNewFile(timeValue)
	}

	return err
}

//y年m月d日h小时
func (f *LogFile) getFileName(t time.Time) string {
	var fileName string
	var err error
	var index = f.fileIndex
	for ; err == nil && index < f.MaxFileNum; index++ {
		fileName = f.FilePath + fmt.Sprintf("%s%s%04d%s", f.FileName, t.Format(f.FileDatePattern), index, f.Suffix)
		_, err = os.Lstat(fileName)
	}

	f.fileIndex = index
	return fileName
}

func (f *LogFile) initTimeRef() {
	//strings.Replace("", "y", strconv.Itoa(10), 1)
	if index := strings.Index(f.FileDatePattern, "1"); index >= 0 {
		f.isDailyRefresh = true
	}

	if index := strings.Index(f.FileDatePattern, "2"); index >= 0 {
		f.isMonthRefresh = true
	}

	if index := strings.Index(f.FileDatePattern, "2006"); index >= 0 {
		f.isYearRefresh = true
	}

	if index := strings.Index(f.FileDatePattern, "15"); index >= 0 {
		f.isHourRefresh = true
	}

}

func (f *LogFile) canSwitchFile(year int, moth int, day int, hour int) bool {
	return (f.isHourRefresh && hour != f.fileHour) ||
		(f.isDailyRefresh && day != f.fileDay) ||
		(f.isMonthRefresh && moth != f.fileMonth) ||
		(f.isYearRefresh && year != f.fileYear) ||
		(f.curLines > f.MaxLines) ||
		(f.curFileSize > f.MaxFileSize)
}

func (f *LogFile) switchNewFile(t time.Time) error {
	f.fileWriteNow.Close()
	newFileName := f.getFileName(t)
	os.Rename(f.filePathName, newFileName)
	return f.createFile()
}

func (f *LogFile) getLines() int {
	fd, err := os.Open(f.filePathName)
	if err != nil {
		return 0
	}
	defer fd.Close()

	buf := make([]byte, 32768) // 32k
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := fd.Read(buf)
		if err != nil && err != io.EOF {
			return count
		}
		count += bytes.Count(buf[:c], lineSep)
		if err == io.EOF {
			break
		}
	}

	return count
}
func init() {
	RegisterAppender(AppenderFile, createLog)
}
