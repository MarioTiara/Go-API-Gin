package repository

import (
	"github.com/MarioTiara/Go-API-Gin/entity"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Book, error)
	findByID(ID int) (entity.Book, error)
	Create(book entity.Book) (entity.Book, error)
}

type repository struct {
	db *gorm.DB
}

// findByID implements Repository
func (r *repository) findByID(ID int) (entity.Book, error) {
	var book entity.Book
	err := r.db.Find(&book, ID).Error

	return book, err
}

func NewBookRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Book, error) {
	var books []entity.Book
	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) Create(book entity.Book) (entity.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) DeleteById(Id int) error {
	err := r.db.Delete(&entity.Book{}, Id).Error
	return err
}

func (r *repository) Update(book entity.Book) error {
	err := r.db.Model(&book).Updates(entity.Book{
		Code:    book.Code,
		Title:   book.Title,
		Author:  book.Author,
		Page:    book.Page,
		Release: book.Release,
	}).Error

	return err
}
