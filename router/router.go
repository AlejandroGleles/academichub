package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Create a new Gin router
	router := gin.Default()

	// Configure CORS
	PermissionCors(router)

	// Initialize routes
	initializeRoutes(router)

	// Run the server
	router.Run(":8080")
}
