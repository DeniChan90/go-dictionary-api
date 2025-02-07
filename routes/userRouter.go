package routes

import (
	controller "github.com/DeniChan90/go-dictionary-api/controllers"
	"github.com/DeniChan90/go-dictionary-api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.RouterGroup) {
	userRouter := incomingRoutes.Group("/users")
	{
		userRouter.Use(middleware.CORSMiddleware())
		userRouter.Use(middleware.Authenticate())
		userRouter.GET("/", controller.GetUsers())
		userRouter.GET("/:user_id", controller.GetUser())
		userRouter.POST("/:user_id/settings", controller.CreateUserSettings())
		userRouter.GET("/:user_id/settings", controller.GetUserSettings())
	}
}
