package routes

import(
    "github.com/gin-gonic/gin"
    controller "github.com/DeniChan90/go-dictionary-api/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
    incomingRoutes.POST("/auth/signup", controller.Signup())
    incomingRoutes.POST("/auth/login", controller.Login())
}
