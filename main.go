package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-api/controllers"
	"go-api/db"
	_ "go-api/docs"
	"go-api/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Book API
// @version         1.0
// @description     API for managing books
// @BasePath /

func main() {
	db.RunMigrations() // Run database migrations (db/migration.go)

	db.InitDatabase() // Initialize database connection (db/init.go)

	// Initialize router
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
