package main

import (
	"fmt"
	"net"

	"github.com/chaithanyaKS/protohacker/servers"
)

func main() {
	tcpListener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("listening error", err)
		return
	}
	defer tcpListener.Close()
	for {

		conn, err := tcpListener.Accept()
		if err != nil {
			fmt.Println("Error while accepting connection", conn)
			return
		}
		go servers.HandleConnection(conn)

	}

}
