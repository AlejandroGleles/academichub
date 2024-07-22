package handler

import (
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func ShowTurmaHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	turma := schemas.Turma{}
	if err := db.First(&turma, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "turma nao encontrado")
		return
	}
	sendSuccess(ctx, "visualizar turma", turma)
}
