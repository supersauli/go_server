package test

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/trace"
	"test/logs"
	"time"
)

func main333333() {
	//开启强大的分析器
	go pprof()
	var logM = logs.LogManage{}
	logM.Init()
	logM.AddAppender(logs.AppenderFile, `{"MaxLines":1000,"FileName":"xxxxxxx"}`)
	for i := 0; i < 10000; i++ {
		msg := "index:%d\n"
		logM.Log(0, msg, i)
	}

	//以下是运行测试(也可以贴你自己的)代码
	// var c sync.Map
	// for i := 0; i < 100; i++ {
	// 	time.Sleep(time.Second * 1)
	// 	go func() {
	// 		for j := 0; j < 1000000; j++ {
	// 			time.Sleep(time.Millisecond * 20)
	// 			c.Store(fmt.Sprintf("%d", j), j)
	// 			fmt.Println(c.Load(fmt.Sprintf("%d", j)))

	// 		}
	// 	}()
	// }
	// time.Sleep(time.Second * 20)
	time.Sleep(time.Second * 20000)
	fmt.Scan()
}

//运行pprof分析器
func pprof() {
	go func() {
		//关闭GC
		debug.SetGCPercent(-1)

		//运行trace
		http.HandleFunc("/start", traces)
		//停止trace
		http.HandleFunc("/stop", traceStop)
		//手动GC
		http.HandleFunc("/gc", gc)
		//网站开始监听
		http.ListenAndServe(":6060", nil)
	}()
}

//手动GC
func gc(w http.ResponseWriter, r *http.Request) {
	runtime.GC()
	w.Write([]byte("StartGC"))
}

//运行trace
func traces(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	w.Write([]byte("TrancStart"))
	fmt.Println("StartTrancs")
}

//停止trace
func traceStop(w http.ResponseWriter, r *http.Request) {
	trace.Stop()
	w.Write([]byte("TrancStop"))
	fmt.Println("StopTrancs")
}
