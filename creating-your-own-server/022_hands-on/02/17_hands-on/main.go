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

	body := "<html>\n<body>"
	body += "<h1>HOLY COW THIS IS LOW LEVEL</h1>"
	body += "\n"
	body += "<div>" + requestMethod + "</div>"
	body += "\n"
	body += "<div>" + requestURI + "</div>"
	body += "\n"

	if requestURI == "/" {
		body += "<div>You are at Home</div>"
		body += "\n"
	}

	if requestURI == "/apply" && requestMethod == "GET" {
		body += "<div>You retrieved from /apply.</div>"
		body += "\n"
	}

	if requestURI == "/apply" && requestMethod == "POST" {
		body += "<div>You posted to /apply.</div>"
		body += "\n"
	}

	body += "</body>\n</html>"

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")

	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))

	fmt.Fprint(c, "Content-Type: text/html\r\n")

	io.WriteString(c, "\r\n")

	io.WriteString(c, body)
}
