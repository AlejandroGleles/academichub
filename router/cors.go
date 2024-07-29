package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func PermissionCors(router *gin.Engine) {

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://127.0.0.1:5500"}, // Adicione outras origens permitidas, se necessário
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))
}
