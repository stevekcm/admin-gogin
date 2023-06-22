package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/user")
	routes.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  200,
		})
	})

}
