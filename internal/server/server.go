package server

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Server struct {
	IP       string
	Port     int
	FilePath string
}

var file_path string

func Start(conf Server) *net.UDPConn {
	address, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", conf.IP, conf.Port))
	if err != nil {
		fmt.Printf("Failed to resolve tcp address: %s\n", err)
		return nil
	}

	fmt.Printf("Starting on %q\n", address.String())

	connection, err := net.ListenUDP("udp", address)
	if err != nil {
		fmt.Printf("Server failed to listen on %q: %s\n", address, err)
		return nil
	}

	fmt.Printf("Server listening on %q\n", address)

	file_path = conf.FilePath

	return connection
}

func Listening(conn *net.UDPConn) {
	for {
		message := make([]byte, 1024)
		_, client, err := conn.ReadFromUDP(message)
		if err != nil {
			fmt.Printf("failed to read from remote client: %s\n", err)
		}

		client_conn, err := net.Dial(client.Zone, fmt.Sprintf("%s:%d", client.IP, client.Port))
		if err != nil {
			fmt.Printf("Failed to dial client: %s\n", err)
			return
		}

		file, err := os.Open(file_path)
		if err != nil {
			fmt.Printf("failed to read file: %s\n", err)
		}
		defer file.Close()

		for {
			buff := make([]byte, 1024)
			_, err = io.ReadFull(file, buff)
			if err == io.EOF {
				fmt.Printf("File ended\n")
				break
			}

			if err != nil && err != io.EOF {
				fmt.Printf("Failed to read file: %s\n", err)
				return
			}

			client_conn.Write(buff)
		}

	}
}
