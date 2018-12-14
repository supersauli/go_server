package main

import (
	"fmt"
	"net"
	"time"
)

func NewCon(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		size, _ := conn.Read(buf)
		msg := string(buf[:size])
		fmt.Println("read size = ", size, msg)
	}
}

func TestSprint() {
	format := "2/1/2006"
	r := fmt.Sprintf("%s.%04d%s", time.Now().Format(format), 10, ".log")
	fmt.Println(r)
}

var g_array = [5]int{5, 4, 3, 2, 1}

func TestArr() {

	var a = g_array
	var b = g_array

	fmt.Printf("%p\n", &g_array[0])
	fmt.Printf("%p\n", &a[0])
	fmt.Printf("%p\n", &b[0])
	b[0] = 99
	fmt.Printf("%d\n", g_array[0])
	fmt.Printf("%d\n", a[0])
	fmt.Printf("%d\n", b[0])
}

func main() {

	TestArr()
	//TestSprint()
	return
	//	ch := make(map[string]string)
	//ch["sss"] = "1111"
	//i, j := ch["sss"]
	//fmt.Println(i)
	//fmt.Println(j)
	blue := string([]byte{27, 91, 52, 50, 109})
	fmt.Printf("%s %s", blue, "xxxxx")
	//fmt.Println(f, "xxxxx")
	fmt.Println("11111111111")

	listen, err := net.Listen("tcp", ":2555")
	if err != nil {
		fmt.Println("listen error", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error", err)
			break
		}

		go NewCon(conn)
	}
}
