package handler

import (
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func ListAtividadeHandler(ctx *gin.Context) {
	atividade := []schemas.Atividade{}

	if err := db.Find(&atividade).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "erro ao listar atividades")
		return
	}
	sendSuccess(ctx, "lista de atividade", atividade)
}
