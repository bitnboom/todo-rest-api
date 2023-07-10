package main

import (
	"fmt"
	"log"
	"todo-rest-api/controller"
	"todo-rest-api/database"
	"todo-rest-api/middleware"
	"todo-rest-api/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    loadEnv()
    loadDatabase()
}

func loadEnv() {
    err := godotenv.Load(".env.local")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func loadDatabase() {
    database.Connect()
    database.Database.AutoMigrate(&model.User{})
    database.Database.AutoMigrate(&model.Todo{})
}
