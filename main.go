package main

import (
	"github.com/avinashb98/munshee/application"
	"github.com/avinashb98/munshee/server"
)

func main() {
	app := application.Get()
	server.StartHTTP(app)
}
