package main

import (
	"fmt"
	"net"
	"strings"
)

func HandlerConnect(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr()
	fmt.Println(addr, "客户端成功连接")

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("服务器检测到客户端已关闭， 断开连接!!!")
			return
		}
		if err != nil {
			fmt.Println("listener.Read error", err)
			return
		}
		fmt.Println("服务器读到数据：", string(buf[:n]))

		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Listen error", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器等待服务端连接：")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error", err)
			return
		}

		go HandlerConnect(conn)
	}
}
