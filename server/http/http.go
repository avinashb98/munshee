package http

import (
	"github.com/avinashb98/munshee/application"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var router = gin.Default()

func StartServer(application *application.Application) {
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "OK",
			"test":   os.Getenv("TEST"),
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

	txnHandler := NewTxnHandler(application.Services.Txn)
	txnRoutes := router.Group("/api/v1/txns")
	{
		txnRoutes.POST("/", txnHandler.CreateTxn)
		txnRoutes.GET("/:username/:id", txnHandler.Get)
		txnRoutes.GET("/:username", txnHandler.GetAll)
		txnRoutes.PUT("/:username/:id/tags", txnHandler.UpdateTags)
	}

	tagHandler := NewTagHandler(application.Services.Tag)
	tagRoutes := router.Group("/api/v1/tags")
	{
		tagRoutes.GET("/", tagHandler.GetAll)
	}
	err := router.Run("0.0.0.0:" + application.Config.Server.Port)
	if err != nil {
		log.Panic(err)
	}
}
