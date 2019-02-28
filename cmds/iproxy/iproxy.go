package main

import (
	"fmt"
	"net"
	"syscall"
)

const (
	SO_ORIGINAL_DST = 80
)

func handleConn(conn net.Conn) {
	tcpConn := conn.(*net.TCPConn)
	tcpConnFile, err := tcpConn.File()
	if err != nil {
		fmt.Println("get tcp conn file failed:", err)
		return
	}

	addr, err := syscall.GetsockoptIPv6Mreq(int(tcpConnFile.Fd()), syscall.IPPROTO_IP, SO_ORIGINAL_DST)
	if err != nil {
		fmt.Println("GetsockoptIPv6Mreq failed:", err)
		return
	}
	fmt.Println("real addr:", addr)

	hostPort := fmt.Sprintf("%d.%d.%d.%d:%d",
		addr.Multiaddr[4],
		addr.Multiaddr[5],
		addr.Multiaddr[6],
		addr.Multiaddr[7],
		uint16(addr.Multiaddr[2])<<8+uint16(addr.Multiaddr[3]))
	fmt.Println("hostPort:", hostPort)

	fmt.Println("start close conn...")
	_ = tcpConnFile.Close()
	_ = tcpConn.Close()
}

func main() {

	l, err := net.Listen("tcp", ":1190")
	if err != nil {
		fmt.Println("net.Listen failed:", err)
		return
	}

	fmt.Println("Listen success, waiting accept")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Accept connection failed:", err)
			return
		}
		go handleConn(conn)
	}

	fmt.Println("Hello, world")
}
