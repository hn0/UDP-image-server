package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"gopkg.in/gographics/imagick.v3/imagick"
	"log"
	"math"
	"net"
	"netimage"
	"os"
)

var ROOT string
var DEBUG bool = true
var buffsz int = 2048

func response(conn *net.UDPConn, sender *net.UDPAddr, data []byte) {
	i := 0
	for {
		j := int(math.Min(float64(buffsz), float64(len(data)-i)))
		// log.Println("buff iter", i, j, len(data))
		conn.WriteTo(data[i:i+j], sender)
		i = i + j
		if i >= len(data) {
			break
		}
	}
}

func recive(conn *net.UDPConn) {
	imagick.Initialize()
	defer imagick.Terminate()

	buff := make([]byte, buffsz)
	var empty []byte = []byte("-1")

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
			img.File = ROOT + string(buff[8:n])

			if blob := img.Resize(); blob != nil {
				go response(conn, sender, blob)
			} else {
				log.Println("Image resize fail", img.File)
				response(conn, sender, empty)
			}

		}
	}
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please provide root path from which this app can serve images")
		os.Exit(1)
	}

	ROOT = os.Args[1]

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
