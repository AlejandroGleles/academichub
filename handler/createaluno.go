package handler

import (
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func CreateAlunoHandler(ctx *gin.Context) {
	request := CreateAlunoRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	aluno := schemas.Aluno{
		Nome:      request.Nome,
		Matricula: float64(request.Matricula),
		Turma:     request.Turma,
	}

	if err := db.Create(&aluno).Error; err != nil {
		logger.Errorf("error creating openig: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating aluno on database")
		return
	}
	sendSuccess(ctx, "create-proffesor", aluno)

}
