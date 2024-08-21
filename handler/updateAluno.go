package handler

import (
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateAlunoHandler(ctx *gin.Context) {
	request := UpdateAlunoRequest{}

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
	aluno := schemas.Aluno{}
	if err := db.First(&aluno, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "aluno nao encontrado")
		return
	}
	//update aluno
	if request.Nome != "" {
		aluno.Nome = request.Nome
	}
	if request.Matricula >= 0 {
		aluno.Matricula = float64(request.Matricula)
	}
	if request.Turma != "" {
		aluno.Turma = request.Turma
	}
	//save aluno
	if err := db.Save(&aluno).Error; err != nil {
		logger.Errorf("error updating aluno: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating ")
		return
	}
	sendSuccess(ctx, "update aluno", aluno)

}
