package http

import (
	"github.com/avinashb98/munshee/service"
	"github.com/gin-gonic/gin"
)

type TagHandler interface {
	GetAll(c *gin.Context)
}

type tagHandler struct {
	tagService service.Tag
}

func (t tagHandler) GetAll(c *gin.Context) {
	tags, err := t.tagService.GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"tags": tags,
	})
}

func NewTagHandler(tagService service.Tag) TagHandler {
	return &tagHandler{
		tagService: tagService,
	}
}
