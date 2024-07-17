package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpenigHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Post Openig",
	})
}
