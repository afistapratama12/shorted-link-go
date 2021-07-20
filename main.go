package main

import (
	"fmt"
	"shortedLink/shortener"
)

var (
// shortService = service.NewService(service.Memmory)
// handler      = service.NewShortHandler(shortService)
)

func main() {
	// r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Welcome to the URL shortener APi",
	// 	})
	// })

	// r.GET("/:short_link", handler.RedirectLongLink)
	// r.POST("/", handler.CreateNewShortLink)
	// // r.PUT("/:short_link_id")

	// r.Run()

	short := shortener.GenerateShortLink("https://www.google.com/search?q=afista+pratama&rlz=1C1GCEA_enID937ID937&oq=afista+pratama&aqs=chrome..69i57j46i13i175i199j0i8i13i30j69i60l3.2077j1j15&sourceid=chrome&ie=UTF-8", "")
	short2 := shortener.GenerateShortLink("https://www.google.com/search?q=impact+byte&rlz=1C1GCEA_enID937ID937&oq=impact+byte&aqs=chrome..69i57j35i39j0j0i22i30l7.1454j1j15&sourceid=chrome&ie=UTF-8", "")
	short3 := shortener.GenerateShortLink("https://www.linkedin.com/in/afistapratama", "")

	fmt.Println(short)
	fmt.Println(short2)
	fmt.Println(short3)
}
