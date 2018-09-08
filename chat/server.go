package main

import (
	"fmt"
	"net"
)

// 简易聊天室

// 目的 ： 
	// 1、上线通知
	// 2、消息群发
	// 3、下线通知 sys.singal
	// 4、用户列表 list
	// 5、rename


// 实现：
//  1、每个客户端各跑一个接受和发送的goruntine
// 	2、发到消息中心，群发



var clients = map[string]net.Conn


func main() {

	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		PrintError(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			PrintError(err)
			continue
		} else {
			go handleConnect(conn)
		}
	}

}

func PrintError(err error) {
	fmt.Println("err : ", err.Error())
}

func handleConnect(conn net.Conn) {
	defer conn.Close()
	fmt.Println("conn ", conn, conn.RemoteAddr())
}
