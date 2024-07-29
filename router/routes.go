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
		v1.GET("/professor", handler.ShowProfessorHandler)
		v1.POST("/professor", handler.CreateProfessorHandler)
		v1.DELETE("/professor", handler.DeleteProfessorHandler)
		v1.PUT("/professor", handler.UpdateProfessorHandler)
		v1.GET("/professores", handler.ListProfessorHandler)

	}

	{
		//routs aluno
		v1.GET("/aluno", handler.ShowAlunoHandler)
		v1.POST("/aluno", handler.CreateAlunoHandler)
		v1.DELETE("/aluno", handler.DeleteAlunoHandler)
		v1.PUT("/aluno", handler.UpdateAlunoHandler)
		v1.GET("/alunos", handler.ListAlunorHandler)

	}

	{
		//routs turma
		v1.GET("/turma", handler.ShowTurmaHandler)
		v1.POST("/turma", handler.CreateTurmaHandler)
		v1.DELETE("/turma", handler.DeleteAlunoHandler)
		v1.PUT("/turma", handler.UpdateTurmaHandler)
		v1.GET("/turmas", handler.ListTurmaHandler)

	}

	{
		//routs atividade
		v1.GET("/atividade", handler.ShowAtividadeHandler)
		v1.POST("/atividade", handler.CreateAtividadeHandler)
		v1.DELETE("/atividade", handler.DeleteAtividadeHandler)
		v1.PUT("/atividade", handler.UpdateAtividadeHandler)
		v1.GET("/atividades", handler.ListAtividadeHandler)

	}
}
