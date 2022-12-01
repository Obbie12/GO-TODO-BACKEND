package main

import (
	"fmt"
	"log"
	"os"

	"go-todo-backend/database"
	"go-todo-backend/handlers"
	"go-todo-backend/models"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.Todo{})
}

func serveApplication() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Env var 'PORT' must be set")
	}

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.GET("/todos", handlers.GetAllTodosHandler)
	e.POST("/todos", handlers.CreateTodoHandler)

	e.GET("/todos/:id", handlers.GetTodoHandler)
	e.DELETE("/todos/:id", handlers.DeleteTodoHandler)
	e.PATCH("/todos/:id", handlers.UpdateTodoHandler)

	e.Logger.Fatal(e.Start(":" + port))
	fmt.Printf("Server running on port " + port)
}
