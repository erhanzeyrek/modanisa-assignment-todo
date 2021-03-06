package app

import (
	"fmt"
	"log"
	"os"
	"todo-api/domain"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print(".env file found")
	}
}

func StartApp() {

	dbdriver := os.Getenv("DBDRIVER")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	database := os.Getenv("DATABASE")
	port := os.Getenv("PORT")

	domain.TodoRepo.Initialize(dbdriver, username, password, port, host, database)
	fmt.Println("DATABASE STARTED")

	routes()

	router.Run(":9090")
}