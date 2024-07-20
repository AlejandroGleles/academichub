package handler

import (
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpenigHandler(ctx *gin.Context) {
	request := UpdateOpenigRequest{}

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
	professor := schemas.Professor{}
	if err := db.First(&professor, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "professor not found")
		return
	}
	//update professor
	if request.Nome != "" {
		professor.Nome = request.Nome
	}
	if request.Email != "" {
		professor.Email = request.Email
	}
	if request.CPF != "" {
		professor.CPF = request.CPF
	}
	//save profesor
	if err := db.Save(&professor).Error; err != nil {
		logger.Errorf("error updating professor: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating professor")
		return
	}
	sendSuccess(ctx, "update professor", professor)

}
