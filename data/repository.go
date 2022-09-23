package data

import (
	"github.com/MarioTiara/Go-API-Gin/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.Book, error)
	findByID(ID int) (model.Book, error)
	Create(book model.Book) (model.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]model.Book, error) {
	var books []model.Book
	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) FindByID(id int) (model.Book, error) {
	var book model.Book
	err := r.db.Find(&book, id).Error

	return book, err
}

func (r *repository) Create(book model.Book) (model.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) DeleteById(Id int) error {
	err := r.db.Delete(&model.Book{}, Id).Error
	return err
}
