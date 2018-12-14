package logs

import (
	"testing"
)

func BenchmarkFileLo(t *testing.B) {
	//var log Logger = CreateLog()
	var logM = LogManage{}
	logM.Init()
	logM.AddAppender(AppenderFile, `{"MaxLines":1000,"FileName":"xxxxxxx"}`)
	for i := 0; i < t.N; i++ {
		msg := "index111:%d\n"
		logM.Log(0, msg, i)
	}
	t.Log("end")

	var logM1 = LogManage{}
	logM1.Init()
	logM1.AddAppender(AppenderFile, `{"MaxLines":1000,"FileName":"12233"}`)
	for i := 0; i < t.N; i++ {
		msg := "index:%d\n"
		logM1.Log(0, msg, i)
	}



	//fmt.Printf("test log\n")
	//log.Stop()
}

func TestLog(t *testing.T) {
	var logM = LogManage{}
	//reflect.TypeOf(logM)
	logM.Init()
	logM.AddAppender(AppenderFile, `{"MaxLines":1000,"FileName":"xxxxxxx"}`)
	for i := 0; i < 1; i++ {
		msg := "index:%d\n"
		logM.Log(0, msg, i)
	}
}
