package http

import (
	"fmt"
	"github.com/avinashb98/munshee/entity"
	"github.com/avinashb98/munshee/service"
	"github.com/gin-gonic/gin"
)

type AccountHandler interface {
	CreateAccount(c *gin.Context)
	GetAccount(c *gin.Context)
	GetAllAccounts(c *gin.Context)
}

type accountHandler struct {
	accountService service.Account
}

func (a accountHandler) CreateAccount(c *gin.Context) {
	input, err := a.parseCreateAccountInput(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	newAccount, err := a.accountService.CreateAccount(input.Username, input.Name)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "account created successfully",
		"account": newAccount,
	})
}

func (a accountHandler) parseCreateAccountInput(c *gin.Context) (*entity.AccountIn, error) {
	var payload entity.AccountIn
	if err := c.ShouldBindJSON(&payload); err != nil {
		return nil, err
	}
	if payload.Username == "" {
		return nil, fmt.Errorf("empty username")
	}
	if payload.Name == "" {
		return nil, fmt.Errorf("empty account name")
	}
	return &payload, nil
}

func (a accountHandler) GetAccount(c *gin.Context) {
	username := c.Param("username")
	name := c.Param("name")
	account, err := a.accountService.Get(username, name)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"account": account,
	})
}

func (a accountHandler) GetAllAccounts(c *gin.Context) {
	username := c.Param("username")
	accounts, err := a.accountService.GetAll(username)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"accounts": accounts,
	})
}

func NewAccountHandler(accountService service.Account) AccountHandler {
	return &accountHandler{
		accountService: accountService,
	}
}
