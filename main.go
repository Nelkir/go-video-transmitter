package main

import (
	"strings"

	"go-video-transmitter/internal/envs"
	"go-video-transmitter/internal/server"
)

func main() {
	env := envs.GetEnvs()
	mode := strings.ToLower(env.Mode)
	switch mode {
	case "server":
		connection := server.Start(server.Server{
			IP:       env.ServerIP,
			Port:     env.ServerPort,
			FilePath: env.FilePath,
		})
		server.Listening(connection)
	case "client":
		// connection := client.Connect()
	}
}
