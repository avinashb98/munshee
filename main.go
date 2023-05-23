package main

import (
	"github.com/avinashb98/munshee/application"
	"github.com/avinashb98/munshee/server/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	app := application.Get()
	http.StartServer(app)
}
