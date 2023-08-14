package middleware

import (
	"clean-architecture/util"
	"github.com/gin-gonic/gin"
)

// log request
func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header
		// check x-role-id header is not admin
		if header.Get("x-role-id") != util.ADMIN_ROLE {
			c.AbortWithStatusJSON(403, gin.H{
				"message": "Forbidden",
			})
			return
		}
		c.Next()
	}
}
