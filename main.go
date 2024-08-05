package main

import(
    "os"
    "log"
    "github.com/gin-gonic/gin"
    routes "github.com/DeniChan90/go-dictionary-api/routes"
    "github.com/joho/godotenv"
)


func main(){
    err:= godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error with .env file.")
    }
    port:= os.Getenv("PORT")

    if port == "" {
        port = "8000"
    }

    router := gin.New()
    router.Use(gin.Logger())

    routes.AuthRoutes(router)
    routes.UserRoutes(router)

    router.GET("/api", func(c *gin.Context){
        c.JSON(200, gin.H{"success": "API is OK"})
    })

    router.Run(":" + port)
}
