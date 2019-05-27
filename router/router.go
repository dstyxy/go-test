package router

import (
	. "go-test/api"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/index", Index)
	router.GET("/ab", Ab)
	return router
}
