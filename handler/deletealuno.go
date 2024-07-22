package handler

import (
	"fmt"
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteAlunoHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	aluno := schemas.Aluno{}
	//find aluno
	if err := db.First(&aluno, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("aluno com id: %s nao encontrado", id))
		return
	}
	//Delete aluno
	if err := db.Delete(&aluno).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("erro ao deletar aluno com id: %s", id))
		return
	}
	sendSuccess(ctx, "aluno deletado", aluno)
}
