package http

import (
	"fmt"
	"github.com/avinashb98/munshee/entity"
	"github.com/avinashb98/munshee/service"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
}

type userHandler struct {
	userService service.User
}

func (u userHandler) CreateUser(c *gin.Context) {
	input, err := u.parseCreateUserInput(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	newUser, err := u.userService.CreateUser(input.Username, input.Name, input.Email)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "user created successfully",
		"user":    newUser,
	})
}

func (u userHandler) parseCreateUserInput(c *gin.Context) (*entity.UserIn, error) {
	var payload entity.UserIn
	if err := c.ShouldBindJSON(&payload); err != nil {
		return nil, fmt.Errorf("invalid input: %w", err)
	}
	if payload.Username == "" {
		return nil, fmt.Errorf("empty username")
	}
	if payload.Name == "" {
		return nil, fmt.Errorf("empty name")
	}
	if payload.Email == "" {
		return nil, fmt.Errorf("empty email")
	}
	return &payload, nil
}

func (u userHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := u.userService.Get(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"user": user,
	})
}

func NewUserHandler(userService service.User) UserHandler {
	return &userHandler{
		userService: userService,
	}
}
