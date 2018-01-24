package main

import (
	"bytes"
	"encoding/binary"
	"gopkg.in/gographics/imagick.v3/imagick"
	"log"
	"net"
	"netimage"
	"os"
)

var DEBUG bool = true
var buffsz int = 2048

func response(conn *net.UDPConn, sender *net.UDPAddr, data []byte) {
	conn.WriteToUDP(data, sender)
}

func recive(conn *net.UDPConn) {
	imagick.Initialize()
	defer imagick.Terminate()

	buff := make([]byte, buffsz)
	var empty []byte = []byte("-1")
	// php is shitty with binary data
	// empty := make([]byte, reflect.TypeOf(reflect.Int32).Size())
	// binary.PutVarint(empty, -1)

	for {
		n, sender, err := conn.ReadFromUDP(buff)
		if err == nil {
			defer func() {
				response(conn, sender, empty)
				if r := recover(); r != nil {
					if DEBUG {
						log.Println("Invalid request, could not deliver the content!")
					}
					recive(conn)
				}
			}()
			// log.Println("Got: ", n)
			img := netimage.Image{}
			b := bytes.NewReader(buff)
			binary.Read(b, binary.BigEndian, &img.Twidth)
			binary.Read(b, binary.BigEndian, &img.Thight)
			img.File = string(buff[8:n])

			if blob := img.Resize(); blob != nil {
				log.Println(len(blob))
			}

			response(conn, sender, empty)
		}
	}
}

func main() {

	addr, err := net.ResolveUDPAddr("udp", ":20000")
	if err != nil {
		log.Println("Server addr check error:", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Println("Failed, udp port could not be opened!", err)
		os.Exit(1)
	}
	defer conn.Close()
	recive(conn)

}
