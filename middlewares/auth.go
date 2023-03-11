package midlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Anauthorization")
		if tokenString == "TOKEN-COURSE-NET" {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Anauthorization",
				"status":  http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
	}
}
