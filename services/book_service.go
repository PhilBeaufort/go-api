package services

import (
	"errors"
	"go-api/models"
)

var books = []models.Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func GetBooks() []models.Book {
	return books
}

func GetBookByID(id string) (*models.Book, error) {
	for i := range books {
		if books[i].ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func AddBook(newBook models.Book) models.Book {
	books = append(books, newBook)
	return newBook
}

func CheckoutBook(id string) (*models.Book, error) {
	book, err := GetBookByID(id)
	if err != nil {
		return nil, err
	}
	if book.Quantity <= 0 {
		return nil, errors.New("book not available")
	}
	book.Quantity--
	return book, nil
}

func ReturnBook(id string) (*models.Book, error) {
	book, err := GetBookByID(id)
	if err != nil {
		return nil, err
	}
	book.Quantity++
	return book, nil
}
