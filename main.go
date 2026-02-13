package main

import (
	"fmt"
	"urlshortner/handler"
	"urlshortner/store"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go Url Shortener is working",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStorageService()

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprint("failed to start the web server - ERROR: %v", err))
	}

}
