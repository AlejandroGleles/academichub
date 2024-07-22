package handler

import (
	"net/http"

	"github.com/AlejandroGleles/academichub/schemas"
	"github.com/gin-gonic/gin"
)

func ShowProfessorHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	professor := schemas.Professor{}
	if err := db.First(&professor, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "professor nao encontrado")
		return
	}
	sendSuccess(ctx, "visualizar professor", professor)
}
