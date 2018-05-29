package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		scanner := bufio.NewScanner(conn)

		scanner.Scan()

		fmt.Println("Code got here.")

		io.WriteString(conn, "I see you connected.")

		conn.Close()
	}
}
