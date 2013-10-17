package main

import (
	"fmt"
	"lib/user_casts"
	"lib/user_error"
	"net"
)

func server() {
	ln, err := net.Listen("tcp", ":8080")
	user_error.Check_error(err, user_error.FATAL)
	connection, err := ln.Accept()
	fmt.Printf("%v", connection.RemoteAddr().String())
	user_casts.bytes2int()
	//for {
	//	connection, err := ln.Accept()
	//	user_error.Check_error(err, user_error.ERROR)
	//	go worker(connection, protocal.Enter)
	//}
}

func main() {
	server()
}
