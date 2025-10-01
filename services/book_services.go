package services

import (
	"github.com/shanomz7235/bookstore-back/models"
	"github.com/shanomz7235/bookstore-back/repositories"
)

func CreateBook(book *models.Book) error {
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

func UpdateBook( book *models.Book) error {
	return repositories.UpdateBook( book)
}

func DeleteBook(id uint) error {
	return repositories.DeleteBook(id)
}