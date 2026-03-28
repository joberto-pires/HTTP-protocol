package main

import (
	"fmt"
	"keita_http/internal/request"
	"net"
)

func main() {
	neti, err := net.Listen("tcp", ":4040")
	if err != nil {
		fmt.Println("error : %w", err)
	}
	for {
		conn, err := neti.Accept()
		if err != nil {
			neti.Close()
			break
		}
		req, err := request.RequestFromReader(conn)
		if err != nil {
		  fmt.Println("error : %w", err)
		}
		fmt.Println("Request line: ")
		fmt.Printf("- Method: %s\n", req.RequestLine.Method)
		fmt.Printf("- Target: %s\n", req.RequestLine.RequestTarget)
		fmt.Printf("- Version: %s\n", req.RequestLine.HttpVersion)

  }
}















