package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sqmatheus/todo-api/database"
	router "github.com/sqmatheus/todo-api/routes"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERROR: could not load the .env file")
	}

	db := database.Initialize(os.Getenv("DB_DNS"))

	r := gin.Default()

	router.Initialize(r, db)

	err = r.Run()
	if err != nil {
		log.Fatal("ERROR: could not run the server")
	}
}
