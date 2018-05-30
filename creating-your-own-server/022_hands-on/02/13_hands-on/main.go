package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

	var requestMethod string
	var requestURI string
	i := 0

	for scanner.Scan() {
		ln := scanner.Text()

		if i == 0 {
			requestLine := strings.Fields(ln)
			requestMethod = requestLine[0]
			requestURI = requestLine[1]

			fmt.Println(requestMethod, requestURI)
		}

		fmt.Println(ln)

		i++

		if ln == "" {
			break
		}
	}

	body := "Test body"
	body += "\n"
	body += requestMethod
	body += "\n"
	body += requestURI

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")

	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))

	fmt.Fprint(c, "Content-Type: text/plain\r\n")

	io.WriteString(c, "\r\n")

	io.WriteString(c, body)
}
