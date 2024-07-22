package handler

import (
	"fmt"
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteTurmaHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	turma := schemas.Turma{}
	//find aluno
	if err := db.First(&turma, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("turma com id: %s nao encontrado", id))
		return
	}
	//Delete aluno
	if err := db.Delete(&turma).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("erro ao deletar turma com id: %s", id))
		return
	}
	sendSuccess(ctx, "turma deletado", turma)
}
