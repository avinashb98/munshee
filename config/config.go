package config

type Server struct {
	Port string
}

type Config struct {
	Server Server
}

func New() Config {
	return Config{
		Server: Server{
			Port: "8080",
		},
	}
}
