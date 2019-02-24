package socket

import (
	"log"
	"net"
)

type ServerSocket struct {
	_ipinfo string
	_lsock  *net.TCPListener
}

func (sock *ServerSocket) Init(sockType string, addressAndPort string) {
	tcpaddr, err := net.ResolveTCPAddr("tcp4", addressAndPort)
	if err != nil {
		log.Fatal(err)
	}
	sock._lsock, err = net.ListenTCP("tcp", tcpaddr)
	if err != nil {
		log.Fatal(err)
	}
	sock._ipinfo = addressAndPort
}
func (sock *ServerSocket) Do() {
	for {
		conn, err := sock._lsock.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(co net.Conn) {
	defer co.Close()
}

func main()1 {
	var so ServerSocket
	so.Init("", "127.0.0.1:2344")
	so.Do()
}
