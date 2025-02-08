package main

import (
	"context"
	"flag"

	//	"fmt"
	//"log"
	"net/http"
	"os"
	"time"

	constant "github.com/DeniChan90/go-dictionary-api/constants"
	"github.com/DeniChan90/go-dictionary-api/middleware"
	onlineWs "github.com/DeniChan90/go-dictionary-api/online-socket"
	routes "github.com/DeniChan90/go-dictionary-api/routes"
	chatWs "github.com/DeniChan90/go-dictionary-api/web-socket"

	//gtranslate "github.com/gilang-as/google-translate"
	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
)

var addr = flag.String("addr", ":8080", "http service address")

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

	flag.Parse()
	chatHub := chatWs.NewHub()
	onlineHub := onlineWs.NewHub()

	go chatHub.Run()
	go onlineHub.Run()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routerGroup := router.Group("/api")

	router.Use(middleware.CORSMiddleware())

	AddRoutes(routerGroup)

	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "API is OK"})
	})

	router.GET("/api/available-languages", func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()
		c.JSON(http.StatusOK, constant.Languages)
	})

	router.GET("/chat", func(c *gin.Context) {
		chatWs.ServeWs(chatHub, c.Writer, c.Request)
	})

	router.GET("/online", func(c *gin.Context) {
		onlineWs.ServeWs(onlineHub, c.Writer, c.Request)
	})

	router.Run(":" + port)

}
