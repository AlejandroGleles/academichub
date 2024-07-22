package handler

import (
	"fmt"
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteAtividadeHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	atividade := schemas.Atividade{}
	//find atividade
	if err := db.First(&atividade, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("atividade com id: %s nao encontrado", id))
		return
	}
	//Delete atividade
	if err := db.Delete(&atividade).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("erro ao deletar atividade com id: %s", id))
		return
	}
	sendSuccess(ctx, "atividade deletado", atividade)
}
