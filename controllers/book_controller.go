package controllers

import (
	"go-api/db"
	_ "go-api/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBooks godoc
// @Summary      List all books
// @Description  Returns a list of all books in the catalog
// @Tags         books
// @Produce      json
// @Success      200  {array}   db.Book
// @Router       /books [get]
func GetBooks(c *gin.Context) {
	books, err := db.QueriesInstance.ListBooks(c.Request.Context())

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving books."})
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

// GetbookById godoc
// @Summary      Get a book by ID
// @Description  Returns a single book by ID
// @Tags         books
// @Produce      json
// @Param        id   path      string  true  "Book ID"
// @Success      200  {object}  db.Book
// @Failure      404  {object}  map[string]string
// @Router       /books/{id} [get]
func GetBookByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID"})
		return
	}

	book, err := db.QueriesInstance.GetBook(c.Request.Context(), int32(id))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
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
// @Param        book  body      dto.CreateBookRequest  true  "Book to add"
// @Success      201   {object}  db.Book
// @Router       /books [post]
func CreateBook(c *gin.Context) {
	var newBook db.CreateBookParams

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book data"})
		return
	}

	createdBook, err := db.QueriesInstance.CreateBook(c.Request.Context(), newBook)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Could not create book"})
		return
	}

	c.IndentedJSON(http.StatusCreated, createdBook)
}

// CheckoutBook godoc
// @Summary      Checkout a book
// @Description  Decrease the quantity of the specified book by 1
// @Tags         books
// @Produce      json
// @Param        id   query     string  true  "Book ID"
// @Success      200  {object}  db.Book
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /checkout [patch]
func CheckoutBook(c *gin.Context) {
	idQuery, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	id, err := strconv.Atoi(idQuery)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID"})
		return
	}

	book, err := db.QueriesInstance.CheckoutBook(c.Request.Context(), int32(id))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to checkout book"})
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
// @Success      200  {object}  db.Book
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /return [patch]
func ReturnBook(c *gin.Context) {
	idQuery, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	id, err := strconv.Atoi(idQuery)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID"})
		return
	}

	book, err := db.QueriesInstance.ReturnBook(c.Request.Context(), int32(id))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to return book"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}
