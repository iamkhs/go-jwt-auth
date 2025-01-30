package main

import (
	"github.com/gin-gonic/gin"
	"go-jwt-auth/config"
	"go-jwt-auth/models"
	"go-jwt-auth/routes"
	"log"
)

func main() {
	config.ConnectDatabase()
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Something went wrong!", err)
		return
	}

	router := gin.Default()
	routes.SetupRoutes(router)
	err = router.Run("localhost:8080")
	if err != nil {
		log.Fatal("error couldn't start")
		return
	}
}
