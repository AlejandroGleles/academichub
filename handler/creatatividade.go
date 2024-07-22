package handler

import (
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func CreateAtividadeHandler(ctx *gin.Context) {
	request := CreateAtividadeRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	atividade := schemas.Atividade{
		Turma: request.Turma,
		Valor: request.Valor,
		Data:  request.Data,
	}

	if err := db.Create(&atividade).Error; err != nil {
		logger.Errorf("error creating atividade: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating aluno on database")
		return
	}
	sendSuccess(ctx, "create-Atividade", atividade)

}
