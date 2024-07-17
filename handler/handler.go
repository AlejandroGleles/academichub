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
func ShowOpenigHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Openig",
	})
}
func DeleteOpenigHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Openig",
	})
}
func UpdateOpenigHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Openig",
	})
}
func ListOpenigHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Openig",
	})
}
