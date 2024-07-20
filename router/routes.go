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

		v1.GET("/opening", handler.ShowOpenigHandler)
		v1.POST("/opening", handler.CreateOpenigHandler)
		v1.DELETE("/opening", handler.DeleteOpenigHandler)
		v1.PUT("/opening", handler.UpdateOpenigHandler)
		v1.GET("/openings", handler.ListOpenigHandler)

	}
}
