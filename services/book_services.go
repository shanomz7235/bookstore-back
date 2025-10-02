package services

import (
	"github.com/shanomz7235/bookstore-back/models"
	"github.com/shanomz7235/bookstore-back/repositories"
	"errors"
)

func CreateBook(book *models.Book) error {
	if book.Title == "" || book.Author == "" || book.Price < 0.0 {
		print(book.Price)
		return errors.New("missing or invalid required fields (Title, Author, Price > 0, Stock > 0)")
	}
	print(book.Price)
	return repositories.CreateBook(book)
}

func GetBooks() ( []models.Book, error) {
	return repositories.GetBooks()
}

func GetBook(id uint) (*models.Book, error) {
	book, err := repositories.GetBook(id)
    if err != nil {
        return nil, err
    }
    return book, nil
}

func UpdateBook(id uint, newBook *models.BookUpdate) error {
	book, err := repositories.GetBook(id)
	if err != nil {
		return err
	}

	if newBook.Title == nil && newBook.Author == nil && newBook.Price == nil && newBook.Stock == nil {
		return errors.New("invalid book information")
	}
	return repositories.UpdateBook( book, newBook)
}

func DeleteBook(id uint) error {
	book, err := repositories.GetBook(id)
	if err != nil {
		return err
	}
	return repositories.DeleteBook(id, book)
}