package main

import (
	"shortedLink/config"
	"shortedLink/service"

	"github.com/gin-gonic/gin"
)

// initial variabel in package service
var (
	db           = config.ConnectionDB()
	shortRepo    = service.NewRepository(db)
	shortService = service.NewService(shortRepo)
	handler      = service.NewShortHandler(shortService)
)

func main() {
	// initial server using gin-gonic package
	r := gin.Default()

	// migrate table in dataabse if table not exist
	db.AutoMigrate(&service.ShortedLink{})

	// testing server routes for root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL shortener APi",
		})
	})

	r.GET("/:short_link", handler.RedirectLongLink)
	r.POST("/", handler.CreateNewShortLink)
	r.PUT("/:short_link_id", handler.UpdateShortLink)

	// running on default port :8080
	r.Run()
}
