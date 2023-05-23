package config

import "os"

type Server struct {
	Port string
}

type Config struct {
	Server Server
}

func New() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return Config{
		Server: Server{
			Port: port,
		},
	}
}
