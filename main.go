package main

import (
	"go-test/router"
)

func main() {
	router := router.CreateRouter()
	// router := gin.Default()
	// router.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{"mes": "asd"})
	// })
	// router.GET("/a", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"kj": "aisuhd",
	// 	})
	// })
	// browser.Open("http://127.0.0.1:8080")
	router.Run(":9000")
}
