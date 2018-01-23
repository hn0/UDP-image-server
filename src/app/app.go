package main

import (
	"fmt"
	"net"
	"os"
)

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

	buf := make([]byte, 1024)

	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error in receiving bytes", err)
		} else {
			fmt.Println("Got the: " + string(buf[0:n]))
		}
	}

	fmt.Println("ready to start coding!", addr)
}
