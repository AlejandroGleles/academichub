package handler

import (
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func ShowAtividadeHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	atividade := schemas.Atividade{}
	if err := db.First(&atividade, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "atividade nao encontrada")
		return
	}
	sendSuccess(ctx, "visualizar atividade", atividade)
}
