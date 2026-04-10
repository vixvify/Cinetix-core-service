package main

import (
	"log"
	"os"
	"server/internal/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("CONFIG NOT FOUND")
	}
	database.Connect()
	//jwtSecret := os.Getenv("JWT_SECRET")
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Cookie", "x-from"},
		AllowCredentials: true,
	}))

	//api := r.Group("/api")
	r.Run(":" + PORT + "✅")

}