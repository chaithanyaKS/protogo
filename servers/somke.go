package servers

import (
	"fmt"
	"io"
	"net"
)

func HandlePingConnection(c net.Conn) {
	fmt.Println("Listening on ", c.RemoteAddr().String())
	buffer := make([]byte, 2048)
	defer c.Close()
	for {
		_, err := c.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("reading error", err)
			}
			break
		}
	}
	c.Write(buffer)
}
