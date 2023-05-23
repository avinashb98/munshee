package http

import (
	"github.com/avinashb98/munshee/application"
	"github.com/gin-gonic/gin"
	"log"
)

var router = gin.Default()

func StartServer(application *application.Application) {
	router.Use(gin.Recovery())
	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "OK",
		})
	})

	userHandler := NewUserHandler(application.Services.User)
	userRoutes := router.Group("/api/v1/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUser)
	}

	accountHandler := NewAccountHandler(application.Services.Account)
	accountRoutes := router.Group("/api/v1/accounts")
	{
		accountRoutes.POST("/", accountHandler.CreateAccount)
		accountRoutes.GET("/:username/:name", accountHandler.GetAccount)
		accountRoutes.GET("/:username", accountHandler.GetAllAccounts)
	}

	err := router.Run(":" + application.Config.Server.Port)
	if err != nil {
		log.Panic(err)
	}
}
