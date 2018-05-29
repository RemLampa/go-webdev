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

		go serve(conn)

		fmt.Println("Code got here.")

		io.WriteString(conn, "I see you connected.")
	}
}

func serve(c net.Conn) {
	defer c.Close()

	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if ln == "" {
			break
		}
	}
}
