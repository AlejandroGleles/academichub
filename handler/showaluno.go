package handler

import (
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func ShowAlunoHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	aluno := schemas.Aluno{}
	if err := db.First(&aluno, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "aluno nao encontrado")
		return
	}
	sendSuccess(ctx, "visualizar aluno", aluno)
}
