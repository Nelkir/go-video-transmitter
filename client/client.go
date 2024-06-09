package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	ip, port := os.Getenv("IP"), os.Getenv("PORT")
	stdout, stderr := bufio.NewWriter(os.Stdout), bufio.NewWriter(os.Stderr)

	address, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		stderr.WriteString(fmt.Sprintf("Failed to resolve address %v: %s\n", address, err))
		stderr.Flush()
		return
	}

	connection, err := net.ListenUDP("udp", address)
	if err != nil {
		stderr.WriteString(fmt.Sprintf("Failed to listen %v: %s\n", address, err))
		stderr.Flush()
		return
	}
	defer connection.Close()

	for {
		message := make([]byte, 1024)
		_, err := connection.Read(message)
		if err != nil {
			stderr.WriteString(fmt.Sprintf("Error reading from connection: %s\n", err))
			stderr.Flush()
		}

		_, err = stdout.Write(message)
		if err != nil {
			stderr.WriteString(fmt.Sprintf("Failed to write to stdout: %s\n", err))
			stderr.Flush()
		}
		stdout.Flush()
	}
}
