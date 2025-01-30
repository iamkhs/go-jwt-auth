package routes

import (
	"github.com/gin-gonic/gin"
	"go-jwt-auth/controller"
	"go-jwt-auth/middleware"
	"net/http"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			username, _ := c.Get("username")
			email, _ := c.Get("email")
			c.JSON(http.StatusOK, gin.H{
				"message":  "This is a protected route",
				"username": username,
				"email":    email,
			})
		})
	}
}
