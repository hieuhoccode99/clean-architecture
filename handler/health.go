package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Health struct {
}

func NewHealth() *Health {
	return &Health{}
}

func (h *Health) CheckHealth(c *gin.Context) {
	fmt.Println("good good")
	c.JSON(200, gin.H{
		"message": "Check health success",
	})
}
