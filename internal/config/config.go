package config

import "os"

type Server struct {
	Port string
}

type Config struct {
	Server
}

func Load() *Config {

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	return &Config{
		Server: Server{
			Port: serverPort,
		},
	}
}
