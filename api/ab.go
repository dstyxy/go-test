package api

import (
	"github.com/gin-gonic/gin"
)

func Ab(ctx *gin.Context) {
	ctx.String(200, "ab")
}
