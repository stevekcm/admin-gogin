package main

import (
	"log"

	"github.com/stevekcm/admin-gogin/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db, err := database.ConnectMongo()

	if err != nil {
		log.Println("Failed to connect to MongoDB")
		panic(err)
	}

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	rdb, err := database.ConnectRedis()

	if err != nil {
		log.Println("Failed to connect to Redis")
		panic(err)
	}

	router.Use(func(c *gin.Context) {
		c.Set("rdb", rdb)
		c.Next()
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  200,
		})
	})

	router.Run(":8080")
}
