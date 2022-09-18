package data

import "github.com/MarioTiara/Go-API-Gin/model"

type Books struct {
	item []model.Book
}

func (b *Books) GetBooks() []model.Book {
	return b.item
}

func (b *Books) AddBooks(input model.Book) {
	b.item = append(b.item, input)
}
