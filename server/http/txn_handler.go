package http

import (
	"fmt"
	"github.com/avinashb98/munshee/entity"
	"github.com/avinashb98/munshee/service"
	"github.com/gin-gonic/gin"
)

type TxnHandler interface {
	CreateTxn(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	UpdateTags(c *gin.Context)
}

type txnHandler struct {
	txnService service.Txn
}

func (t txnHandler) CreateTxn(c *gin.Context) {
	input, err := t.parseCreateTxnInput(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	newTxn, err := t.txnService.CreateTxn(*input)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "transaction created successfully",
		"txn":     newTxn,
	})
}

func (t txnHandler) parseCreateTxnInput(c *gin.Context) (*entity.TxnIn, error) {
	var payload entity.TxnIn
	if err := c.ShouldBindJSON(&payload); err != nil {
		return nil, err
	}
	if payload.Username == "" {
		return nil, fmt.Errorf("empty username")
	}
	if payload.FromAccount == nil {
		return nil, fmt.Errorf("transaction must be from an account")
	}
	if payload.Amount == 0 {
		return nil, fmt.Errorf("empty amount")
	}
	return &payload, nil
}

func (t txnHandler) Get(c *gin.Context) {
	username := c.Param("username")
	id := c.Param("txn_id")
	txn, err := t.txnService.Get(username, id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"txn": txn,
	})
}

func (t txnHandler) GetAll(c *gin.Context) {
	username := c.Param("username")
	txns, err := t.txnService.GetAll(username)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"txns": txns,
	})
}

func (t txnHandler) UpdateTags(c *gin.Context) {
	username := c.Param("username")
	id := c.Param("txn_id")
	tags, err := t.parseUpdateTagsInput(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	txn, err := t.txnService.UpdateTags(username, id, tags...)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"txn": txn,
	})
}

func (t txnHandler) parseUpdateTagsInput(c *gin.Context) ([]string, error) {
	var payload struct {
		Tags []string `json:"tags"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		return nil, err
	}
	if len(payload.Tags) == 0 {
		return nil, fmt.Errorf("empty tags")
	}
	return payload.Tags, nil
}

func NewTxnHandler(txnService service.Txn) TxnHandler {
	return &txnHandler{txnService}
}
