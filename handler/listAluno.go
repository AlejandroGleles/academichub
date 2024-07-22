package handler

import (
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func ListAlunorHandler(ctx *gin.Context) {
	aluno := []schemas.Aluno{}

	if err := db.Find(&aluno).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "erro ao listar aluno")
		return
	}
	sendSuccess(ctx, "lista de aluno", aluno)
}
