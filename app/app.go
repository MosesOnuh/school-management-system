package app

import (
	"log"
	"os"

	"management-system/domain"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("failed to load .env file")
	}
}

func StartApplication() {

	dbDriver := os.Getenv("DBDRIVER")
	dbName := os.Getenv("MYSQL_DATABASE")
	username := os.Getenv("MYSQL_USER")
	userPassword := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	domain.DbRepo.Initialize(dbDriver, username, userPassword, port, host, dbName)
	routes()

	router.Run(":8080")
}
