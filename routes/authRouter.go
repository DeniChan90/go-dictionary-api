package routes

import (
	controller "github.com/DeniChan90/go-dictionary-api/controllers"
	"github.com/DeniChan90/go-dictionary-api/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.RouterGroup) {
	authRouter := incomingRoutes.Group("/auth")
	{
		authRouter.POST("/signup", controller.Signup())
		authRouter.POST("/login", controller.Login())
	}

	authRouter.Use(middleware.CORSMiddleware())
	//incomingRoutes.POST("/auth/signup", controller.Signup())
	//incomingRoutes.POST("/auth/login", controller.Login())
}
