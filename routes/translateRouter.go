package routes

import (
	controller "github.com/DeniChan90/go-dictionary-api/controllers"
	"github.com/DeniChan90/go-dictionary-api/middleware"
	"github.com/gin-gonic/gin"
)

func TranslateRoutes(incomingRoutes *gin.RouterGroup) {
	translateRouter := incomingRoutes.Group("/translate")
	{
		translateRouter.Use(middleware.Authenticate())
		translateRouter.POST("/", controller.Translate())
		translateRouter.GET("/languages", controller.GetAvailableLanguages())
		translateRouter.POST("/:user_id/translations", controller.SaveTranslations())
		translateRouter.GET("/:user_id/translations", controller.GetUserTanslations())
	}
}
