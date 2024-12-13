package router

import "github.com/gin-gonic/gin"

func Init() {
	// INITIALIZE ROUTER
	router := gin.Default()

	// INITIALIZE ROUTES
	initRoutes(router)

	// RUN THE SERVER
	router.Run()
}
