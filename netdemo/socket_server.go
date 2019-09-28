package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handlerConn(conn)
	}
}

func handlerConn(conn net.Conn) {
	//如果一个goroutine引发的panic没有得到处理，则会使整个程序退出
	defer func() {
		conn.Close()
		err := recover()
		fmt.Println(err)
	}()

	var buf [1024]byte

	for {
		n, err := conn.Read(buf[0:])
		if err == nil || err == io.EOF {
			s := string(buf[0:n])
			if s == "q" {
				panic("client press q!!")
				break
			}
			fmt.Println(s)
			continue
		}
		log.Println(err)
		break
	}

}
