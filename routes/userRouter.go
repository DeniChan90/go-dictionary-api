package routes

import(
    "github.com/gin-gonic/gin"
    controller "github.com/DeniChan90/go-dictionary-api/controllers"
    "github.com/DeniChan90/go-dictionary-api/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine){
    incomingRoutes.Use(middleware.Authenticate())
    incomingRoutes.GET("/users", controller.GetUsers())
    incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
