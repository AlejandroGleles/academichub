package router

import (
	"github.com/AlejandroGleles/academichub/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	//initialize handler
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		//routs professor
		v1.GET("/professor", handler.ShowOpenigHandler)
		v1.POST("/professor", handler.CreateOpenigHandler)
		v1.DELETE("/professor", handler.DeleteOpenigHandler)
		v1.PUT("/professor", handler.UpdateOpenigHandler)
		v1.GET("/professores", handler.ListOpenigHandler)

	}

	{
		//handler aluno
		//v1.GET("/aluno", handler.ShowAlunoHandler)
		v1.POST("/aluno", handler.CreateAlunoHandler)
		//v1.DELETE("/aluno", handler.DeleteAlunoHandler)
		//v1.PUT("/aluno", handler.UpdateAlunoHandler)
		//v1.GET("/alunos", handler.ListAlunoHandler)

	}

	{
		//handle turma
		//v1.GET("/aluno", handler.ShowAlunoHandler)
		v1.POST("/turma", handler.CreateTurmaHandler)
		//v1.DELETE("/aluno", handler.DeleteAlunoHandler)
		//v1.PUT("/aluno", handler.UpdateAlunoHandler)
		//v1.GET("/alunos", handler.ListAlunoHandler)

	}

	{
		//handler atividade
		//v1.GET("/aluno", handler.ShowAlunoHandler)
		v1.POST("/atividade", handler.CreateAtividadeHandler)
		//v1.DELETE("/aluno", handler.DeleteAlunoHandler)
		//v1.PUT("/aluno", handler.UpdateAlunoHandler)
		//v1.GET("/alunos", handler.ListAlunoHandler)

	}
}
