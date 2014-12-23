package main

import (
	"io"
	"log"
	"net"
	"time"
)

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		log.Print("client recv: ", string(buf[0:n]))
	}
}

func main() {
	c, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		log.Panic(err)
	}
	defer c.Close()

	go reader(c)
	for {
		_, err := c.Write([]byte("hi"))
		log.Print("client send: hi")
		if err != nil {
			log.Fatal("write error: ", err)
			break
		}
		time.Sleep(1e9)
	}
}
