package controllers

import (
	"go-api/models"
	"go-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBooks godoc
// @Summary      List all books
// @Description  Returns a list of all books in the catalog
// @Tags         books
// @Produce      json
// @Success      200  {array}   models.Book
// @Router       /books [get]
func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, services.GetBooks())
}

// GetbookById godoc
// @Summary      Get a book by ID
// @Description  Returns a single book by ID
// @Tags         books
// @Produce      json
// @Param        id   path      string  true  "Book ID"
// @Success      200  {object}  models.Book
// @Failure      404  {object}  map[string]string
// @Router       /books/{id} [get]
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := services.GetBookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

// CreateBook godoc
// @Summary      Add a new book
// @Description  Create a new book entry
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      models.Book  true  "Book to add"
// @Success      201   {object}  models.Book
// @Router       /books [post]
func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book data."})
		return
	}
	createdBook := services.AddBook(newBook)
	c.IndentedJSON(http.StatusCreated, createdBook)
}

// CheckoutBook godoc
// @Summary      Checkout a book
// @Description  Decrease the quantity of the specified book by 1
// @Tags         books
// @Produce      json
// @Param        id   query     string  true  "Book ID"
// @Success      200  {object}  models.Book
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /checkout [patch]
func CheckoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	book, err := services.CheckoutBook(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

// ReturnBook godoc
// @Summary      Return a book
// @Description  Increase the quantity of the specified book by 1
// @Tags         books
// @Produce      json
// @Param        id   query     string  true  "Book ID"
// @Success      200  {object}  models.Book
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /return [patch]
func ReturnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	book, err := services.ReturnBook(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}
