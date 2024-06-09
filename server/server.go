package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	ip, port := os.Getenv("IP"), os.Getenv("PORT")
	file_path := os.Getenv("FILE_PATH")

	client_address, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		fmt.Print("Failed to resolve address: %s\n", err)
	}

	client_connection, err := net.DialUDP("udp", nil, client_address)
	if err != nil {
		fmt.Printf("Failed to dial server: %s\n", err)
		return
	}
	defer client_connection.Close()

	file, err := os.Open(file_path)
	if err != nil {
		fmt.Printf("Failed to open file %s: %s\n", file_path, err)
		return
	}

	file_reader := bufio.NewReader(file)

	for {
		buf := make([]byte, 1024)
		i, err := file_reader.Read(buf)
		if err == io.EOF {
			fmt.Printf("File ended\n")
			break
		}
		fmt.Printf("Sending %d bytes\n", i)
		if err != nil && err != io.EOF {
			fmt.Printf("Failed to read file: %s\n", err)
		}
		client_connection.Write(buf)
	}
}
