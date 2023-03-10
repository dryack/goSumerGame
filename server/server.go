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
	"os"
	"strconv"
)

func main() {
	var (
		port = 80
	)

	loadEnv()
	loadDatabase()
	serveApplication(port)
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
	err := database.Database.AutoMigrate(&model.Game{})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func serveApplication(port int) {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/game", controller.AddGame)
	protectedRoutes.POST("/game/play", controller.TakeTurn)
	protectedRoutes.GET("/game", controller.GetAllGames)
	protectedRoutes.POST("/game/delete", controller.DeleteGame)

	portStr := strconv.Itoa(port)
	err := router.Run(":" + portStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Server running on port %s\n", portStr)
}
