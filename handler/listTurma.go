package handler

import (
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func ListTurmaHandler(ctx *gin.Context) {
	turma := []schemas.Turma{}

	if err := db.Find(&turma).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "erro ao listar turmas")
		return
	}
	sendSuccess(ctx, "lista de turmas", turma)
}
