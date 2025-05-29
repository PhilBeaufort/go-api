package routes

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBookByID)
	router.POST("/books", controllers.CreateBook)
	router.PATCH("/checkout", controllers.CheckoutBook)
	router.PATCH("/return", controllers.ReturnBook)
}
