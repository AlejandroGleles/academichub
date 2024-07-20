package handler

import (
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpenigHandler(ctx *gin.Context) {
	professor := []schemas.Professor{}

	if err := db.Find(&professor).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "erro ao listar professores")
		return
	}
	sendSuccess(ctx, "lista de professores", professor)
}
