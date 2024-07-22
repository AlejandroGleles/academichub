package handler

import (
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateTurmaHandler(ctx *gin.Context) {
	request := UpdateTurmaRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
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
	//update turma
	if request.Nome != "" {
		turma.Nome = request.Nome
	}
	if request.Semestre == "" {
		turma.Semestre = request.Semestre
	}
	if request.Ano != "" {
		turma.Ano = request.Ano
	}
	//save turma
	if err := db.Save(&turma).Error; err != nil {
		logger.Errorf("error updating turma: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating ")
		return
	}
	sendSuccess(ctx, "update turma", turma)

}
