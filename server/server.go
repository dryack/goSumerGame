package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"goSumerGame/server/controller"
	"goSumerGame/server/database"
	"goSumerGame/server/middleware"
	"goSumerGame/server/model"
	"log"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
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
	database.Database.AutoMigrate(&model.Game{})
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controller.AddGame)
	protectedRoutes.GET("/entry", controller.GetAllGames)
	protectedRoutes.POST("/entry/delete", controller.DeleteGame)

	router.Run(":80")
	fmt.Println("Server running on port 8000")
}
