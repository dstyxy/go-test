package api

import (
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.String(200, "index")
}
