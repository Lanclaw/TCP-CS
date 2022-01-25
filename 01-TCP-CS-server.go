package main

import (
	"fmt"
	"net"
)

func main() {
	// 指定服务器的通信协议，ip地址，端口号。创建一个用于监听的socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen() error: ", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器等待客户端建立连接...")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept error: ", err)
		return
	}
	defer conn.Close()

	fmt.Println("服务器与客户端成功建立连接")
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read error: ", err)
		return
	}

	conn.Write(buf[:n])
	fmt.Println("服务器读到数据：", n, string(buf[:n]))

}
