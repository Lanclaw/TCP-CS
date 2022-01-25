package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Dial error: ", err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("are you ready?"))

	buf := make([]byte, 4096)

	n, err := conn.Read(buf)

	if err != nil {
		fmt.Println("conn.Read", err)
		return
	}

	fmt.Println("服务器回发数据：", n, string(buf[:n]))
}
