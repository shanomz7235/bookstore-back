package repositories

import (

	"github.com/shanomz7235/bookstore-back/config"
	"github.com/shanomz7235/bookstore-back/models"
	"errors"

)

func CreateBook(book *models.Book) error {
	result := config.DB.Create(book)
	if result.Error != nil {
        return result.Error
    }
    return  nil
}

func GetBooks() ( []models.BookResponse, error) {
	var books []models.BookResponse
	result := config.DB.Model(&models.Book{}).
		Order("id").
		Find(&books)
	if result.Error != nil{
		return nil, result.Error
	}

	return books, nil
}

func GetBook(id uint) (*models.BookResponse, error) {
    var book models.BookResponse
    result := config.DB.Model(&models.Book{}).
        Select("id", "title", "author", "price", "stock").
        Where("id = ?", id).
        First(&book)

    if result.Error != nil {
        return nil, result.Error
    }
    return &book, nil
}


func GetBookEntityByID(id uint) (*models.Book, error) {
    var book models.Book
    result := config.DB.First(&book, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &book, nil
}


func UpdateBook(book *models.Book, newBook *models.BookUpdate) error {


	result := config.DB.Model(&book).Updates(newBook)
	if result.Error != nil {
        return result.Error
    }
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}

    return  nil
}

func DeleteBook(id uint, book *models.Book) error {
	
	result := config.DB.Where("id = ?", id).Delete(&book)
	if result.Error != nil{
		return result.Error
	}
	return nil
}