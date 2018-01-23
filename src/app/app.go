package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

type req struct {
	width  uint32
	height uint32
	fname  []byte
}

var buffsz int = 2048

func response(conn *net.UDPConn, sender *net.UDPAddr) {
	conn.WriteToUDP([]byte("Not implemented yet!"), sender)
}

func recive(conn *net.UDPConn) {
	buff := make([]byte, buffsz)

	for {
		n, sender, err := conn.ReadFromUDP(buff)
		if err == nil {
			defer func() {
				response(conn, sender)
				if r := recover(); r != nil {
					fmt.Println("Invalid request detected!")
					recive(conn)
				}
			}()
			// fmt.Println("Got: ", n)
			req := req{}
			b := bytes.NewReader(buff)
			binary.Read(b, binary.BigEndian, &req.width)
			binary.Read(b, binary.BigEndian, &req.height)
			req.fname = buff[8:n]
			fmt.Println(req)
		}
	}
}

func main() {

	addr, err := net.ResolveUDPAddr("udp", ":20000")
	if err != nil {
		fmt.Println("Server addr check error:", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Failed, udp port could not be opened!", err)
		os.Exit(1)
	}
	defer conn.Close()
	recive(conn)

}
