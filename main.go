package main

import (
	"context"

	constant "github.com/DeniChan90/go-dictionary-api/constants"
	"github.com/DeniChan90/go-dictionary-api/middleware"
	routes "github.com/DeniChan90/go-dictionary-api/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
	//"github.com/joho/godotenv"
)

func AddRoutes(routeGroup *gin.RouterGroup) {
	routes.AuthRoutes(routeGroup)
	routes.UserRoutes(routeGroup)
	routes.TranslateRoutes(routeGroup)
}

func main() {
	//err := godotenv.Load(".env")

	//	if err != nil {
	//		log.Fatal("Error with .env file.")
	//	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routerGroup := router.Group("/api")

	routerGroup.Use(middleware.CORSMiddleware())

	AddRoutes(routerGroup)

	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "API is OK"})
	})

	router.GET("/api/available-languages", func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()
		c.JSON(http.StatusOK, constant.Languages)
	})

	router.Run(":" + port)
}
