package envs

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Envs struct {
	ServerIP   string
	ServerPort int
	ClientIP   string
	ClientPort int
	Mode       string
	FilePath   string
}

func GetEnvs() Envs {
	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		fmt.Printf(
			"Failed to parse server port variable (probably SERVER_PORT is empty): %s\n",
			err,
		)
		serverPort = 8080
	}

	clientPort, err := strconv.Atoi(os.Getenv("CLIENT_PORT"))
	if err != nil {
		fmt.Printf(
			"Failed to parse client port variable (probably SERVER_IP is empty): %s\n",
			err,
		)
		clientPort = 8080
	}

	mode := strings.ToLower(os.Getenv("MODE"))
	if mode != "server" && mode != "client" {
		fmt.Printf("Mode is not supported or empty: %q. Supported: server, client\n", mode)
		mode = "server"
	}

	return Envs{
		ServerIP:   os.Getenv("SERVER_IP"),
		ServerPort: serverPort,
		ClientIP:   os.Getenv("CLIENT_IP"),
		ClientPort: clientPort,
		Mode:       mode,
		FilePath:   os.Getenv("FILE_PATH"),
	}
}
