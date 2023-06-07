package main

import (
	"fmt"
	"github.com/avinashb98/munshee/application"
	"github.com/avinashb98/munshee/server/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}
	app := application.Get()
	http.StartServer(app)
}
