package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/neerajbg/go-gin-auth/database"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error in loading env file.")
	}

	database.ConnectDB()
}

func main() {

	// Close the db connection using defer clause
	sqlDb, err := database.DBConn.DB()

	if err != nil {
		log.Println("Error in getting db conn.")
	}

	defer sqlDb.Close()

	port := os.Getenv("port")

	if port == "" {
		port = "8001"
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "Hello World!",
		})
	})

	log.Fatal(router.Run(":" + port))
}
