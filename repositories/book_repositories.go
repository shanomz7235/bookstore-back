package repositories

import (

	"github.com/shanomz7235/bookstore-back/config"
	"github.com/shanomz7235/bookstore-back/models"
)

func CreateBook(book *models.Book) error {
	result := config.DB.Create(book)
	if result.Error != nil {
        return result.Error
    }
    return  nil
}

func GetBooks() ( []models.Book, error) {
	var books []models.Book
	result := config.DB.Order("id").Find(&books)
	if result.Error != nil{
		return nil, result.Error
	}

	return books, nil
}

func GetBook(id uint) ( *models.Book, error) {
	var book models.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
        return nil, result.Error
    }
    return &book, nil

}

func UpdateBook( book *models.Book) error {
	result := config.DB.Model(&book).Updates(book)
	if result.Error != nil {
        return result.Error
    }
    return  nil
}

func DeleteBook(id uint) error {
	var book models.Book
	result := config.DB.Delete(&book,id)
	if result != nil{
		return result.Error
	}
	return nil
}