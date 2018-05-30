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
	body := "Test body"

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")

	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))

	fmt.Fprint(c, "Content-Type: text/plain\r\n")

	io.WriteString(c, "\r\n")

	io.WriteString(c, body)
}
