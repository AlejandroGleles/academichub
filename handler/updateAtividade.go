package handler

import (
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateAtividadeHandler(ctx *gin.Context) {
	request := UpdateAtividadeRequest{}

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
	atividade := schemas.Atividade{}
	if err := db.First(&atividade, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "atividade nao encontrada")
		return
	}
	//update atividade
	if request.Turma != "" {
		atividade.Turma = request.Turma
	}
	if request.Valor <= 0 {
		atividade.Valor = request.Valor
	}
	if request.Data != "" {
		atividade.Data = request.Data
	}
	//save atividade
	if err := db.Save(&atividade).Error; err != nil {
		logger.Errorf("error updating atividade: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating ")
		return
	}
	sendSuccess(ctx, "update atividade", atividade)

}
