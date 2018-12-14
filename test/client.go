package test

import (
	"fmt"
	"net"
	"time"
)

func main1() {
	conn, err := net.Dial("tcp", "127.0.0.1:2555")
	if err != nil {
		fmt.Println("conn error", err)
		return
	}
	for {
		time.Sleep(time.Second * 1)
		buf := make([]byte, 64)
		buf = []byte("12345")
		conn.Write(buf)
	}
}
