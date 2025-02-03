package main

import (
	"log"
	"os"
	"reflect"

	routes "github.com/DeniChan90/go-dictionary-api/routes"
	//"github.com/biter777/countries"
	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
	"golang.org/x/text/language"
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

	//routes.AuthRoutes(router)
	//routes.UserRoutes(router)

	routerGroup := router.Group("/api")
	AddRoutes(routerGroup)

	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "API is OK"})
	})

	//countryUA := countries.Ukraine

	log.Print(">>>", reflect.TypeOf(language.Vietnamese))
	router.Run(":" + port)
}
