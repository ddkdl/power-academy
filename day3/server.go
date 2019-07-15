package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func handleConnection(c net.Conn) {
	fmt.Println(c.RemoteAddr().String(), "joined the room!")

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')

		if err == io.EOF {
			fmt.Println(c.RemoteAddr().String(), "Disconnected.")
			return
		} else if err != nil {
			log.Fatal(err)
		}

		netData = strings.TrimSuffix(netData, "\n")
		fmt.Println(c.RemoteAddr().String()+":", netData)
	}
	c.Close()
}

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer ln.Close()

	for {
		connection, err := ln.Accept()

		if err != nil {
			log.Fatal(err)
			return
		}

		go handleConnection(connection)
	}
}
