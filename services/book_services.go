package services

import (
	"errors"
	"fmt"
	"github.com/shanomz7235/bookstore-back/models"
	"github.com/shanomz7235/bookstore-back/repositories"
)

func CreateBook(book []models.Book) error {
	for i, book := range book {
		if book.Title == "" || book.Author == "" || book.Price < 0.0 {
			return fmt.Errorf("validation failed for book index %d: missing required fields", i)
		}
	}

	return repositories.CreateBook(book)
}

func GetBooks() ([]models.BookResponse, error) {
	return repositories.GetBooks()
}

func GetBook(id uint) (*models.BookResponse, error) {
	book, err := repositories.GetBook(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func UpdateBook(id uint, newBook *models.BookUpdate) error {
	book, err := repositories.GetBookEntityByID(id)
	if err != nil {
		return err
	}

	if newBook.Title == nil && newBook.Author == nil && newBook.Price == nil && newBook.Stock == nil {
		return errors.New("invalid book information")
	}
	return repositories.UpdateBook(book, newBook)
}

func DeleteBook(id uint) error {
	book, err := repositories.GetBookEntityByID(id)
	if err != nil {
		return err
	}
	return repositories.DeleteBook(id, book)
}
