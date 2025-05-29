package main

import (
	_ "go-api/docs"

	"go-api/controllers"
	"go-api/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Book API
// @version         1.0
// @description     API for managing books
// @host            localhost:8080
// @BasePath        /

func main() {
	router := gin.Default()

	// Redirect "/" to Swagger docs
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	router.GET("/ping", controllers.Ping)
	router.GET("/health", controllers.Health)
	routes.RegisterRoutes(router)

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("🚀 Server starting on port 8080...")
	router.Run(":8080")

}
