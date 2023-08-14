package handler

import (
	"clean-architecture/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Migration struct {
	db *gorm.DB
}

func NewMigration(db *gorm.DB) *Migration {
	return &Migration{
		db: db,
	}
}

func (h *Migration) Migrate(c *gin.Context) {
	if err := h.db.AutoMigrate(&model.User{}); err != nil {
		c.JSON(400, gin.H{
			"message": "Migrate failed",
		})
		return
	}
}
