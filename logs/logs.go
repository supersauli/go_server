package logs

import (
	"fmt"
	"time"
)

const (
	AppenderFile = iota
	AppenderConsole
)
const (
	LogLevelTrace = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

var levelStr = [LogLevelFatal + 1]string{"[Trace] ", "[Info] ", "[Warn] ", "[Error] ", "[Fatal] "}

type Logger interface {
	Init(config string) error
	Log(level int, msg string, data ...interface{})
	Stop()
}

type createLogger func() Logger

var appenderGroup = make(map[int]createLogger)

func RegisterAppender(appender int, create createLogger) bool {
	if _, ok := appenderGroup[appender]; ok {
		fmt.Println("all ready regisiter ", appender)
		return false
	}
	appenderGroup[appender] = create
	return true
}

type LogManage struct {
	appender map[int]Logger
}

func (l *LogManage) AddAppender(a int, config string) bool {
	if _, ok := l.appender[a]; ok {
		fmt.Println("all ready AddAppender  ", a)
		return false
	}
	logfunc, ok := appenderGroup[a]
	if !ok {
		fmt.Println("not regisiter ", a)
		return false
	}
	log := logfunc()
	error := log.Init(config)
	if error != nil {
		fmt.Println("AddAppender init error  ", error.Error())
		return false
	}

	l.appender[a] = log
	return true
}
func (l *LogManage) Log(level int, msg string, data ...interface{}) {
	for _, app := range l.appender {
		app.Log(level, msg, data...)
	}
}
func (l *LogManage) Init() {
	l.appender = make(map[int]Logger)
}

func GetLogHeard(timeValue time.Time, log_level int) string {
	//y, m, d := timeValue.Date()
	h, mi, s := timeValue.Clock()
	ns := timeValue.Nanosecond()
	return fmt.Sprintf("[%d:%d:%d:%d] %s", h, mi, s, ns, levelStr[log_level])
}
