package handler

import (
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func CreateTurmaHandler(ctx *gin.Context) {
	request := CreateTurmaRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	turma := schemas.Turma{
		Nome:      request.Nome,
		Semestre:  request.Semestre,
		Ano:       request.Ano,
		Professor: request.Professor,
	}

	if err := db.Create(&turma).Error; err != nil {
		logger.Errorf("error creating openig: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating aluno on database")
		return
	}
	sendSuccess(ctx, "create-turma", turma)

}
