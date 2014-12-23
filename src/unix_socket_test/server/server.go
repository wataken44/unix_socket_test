package main

import (
	"log"
	"net"
)

func echoServer(c net.Conn) {
	for {
		buf := make([]byte, 512)

		n, err := c.Read(buf)
		if err != nil {
			return
		}
		data := buf[0:n]
		log.Print("server recv: ", string(data))

		_, err = c.Write(data)
		log.Print("server send: ", string(data))
		if err != nil {
			log.Fatal("write error: ", err)
		}
	}
}

func main() {
	ls, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal("listen error: ", err)
	}

	for {
		as, err := ls.Accept()
		if err != nil {
			log.Fatal("accept error: ", err)
		}
		log.Print("server accept")
		go echoServer(as)
	}
}
