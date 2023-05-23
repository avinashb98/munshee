package main

import (
	"github.com/avinashb98/munshee/application"
	"github.com/avinashb98/munshee/server/http"
)

func main() {
	app := application.Get()
	http.StartServer(app)
}
