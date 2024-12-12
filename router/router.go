package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// INITIALIZE ROUTER
	router := gin.Default()

	// INITIALIZE ROUTES
	initializeRoutes(router)

	// RUN THE SERVER
	router.Run()
}
