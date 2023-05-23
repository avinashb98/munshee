package server

import (
	"github.com/avinashb98/munshee/application"
	"github.com/gin-gonic/gin"
	"log"
)

var router = gin.Default()

func StartHTTP(application *application.Application) {
	router.Use(gin.Recovery())
	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "OK",
		})
	})
	err := router.Run(":" + application.Config.Server.Port)
	if err != nil {
		log.Panic(err)
	}
}
