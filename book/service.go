package book

import "time"

type Service interface {
	FindAll() ([]Book, error)
	findByID(ID int) (Book, error)
	Create(BookRequest BookRequest) (Book, error)
	deleteByID(ID int) error
	Update(BookRequest BookRequest) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FinAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.findByID(ID)
	return book, err
}

func (s *service) Create(BookRequest BookRequest) (Book, error) {
	price, _ := BookRequest.Price.Int64()
	rating, _ := BookRequest.Rating.Int64()
	book := Book{
		Title:     BookRequest.Title,
		Price:     int(price),
		Rating:    int(rating),
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	newbook, err := s.repository.Create(book)
	return newbook, err
}

func (s *service) deleteByID(ID int) error {
	err := s.repository.deleteByID(ID)
	return err
}

func (s *service) Update(BookRequest BookRequest) error {
	price, _ := BookRequest.Price.Int64()
	rating, _ := BookRequest.Rating.Int64()
	book := Book{
		Title:     BookRequest.Title,
		Price:     int(price),
		Rating:    int(rating),
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	err := s.repository.Update(book)
	return err
}
