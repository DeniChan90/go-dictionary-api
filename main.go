package main

import (
	"context"
	//	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	constant "github.com/DeniChan90/go-dictionary-api/constants"
	"github.com/DeniChan90/go-dictionary-api/middleware"
	routes "github.com/DeniChan90/go-dictionary-api/routes"
	gtranslate "github.com/gilang-as/google-translate"
	"github.com/gin-gonic/gin"
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

	value := gtranslate.Translate{
		Text: "Hello",
		From: "en",
		To:   "uk",
	}
	translated, err := gtranslate.Translator(value)
	if err != nil {
		panic(err)
	} else {
		//prettyJSON, err := json.MarshalIndent(translated, "", "\t")
		if err != nil {
			panic(err)
		}
		log.Print("Test log >>>>", translated)
	}
	router.Run(":" + port)

}
