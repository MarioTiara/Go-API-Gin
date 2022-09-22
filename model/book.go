package model

type Book struct {
	Id      int     `json:"id" binding:"required"`
	Code    string  `json:"Code" binding:"required"`
	Title   string  `json:"title" binding:"required"`
	Author  string  `json:"author" binding:"required"`
	Page    int     `json:"page"`
	Price   float32 `json:"price"`
	Release int     `json:"release"`
}

type Books struct {
	Item []Book
}

func (b *Books) GetBooks() []Book {
	return b.Item
}

func (b *Books) AddBooks(input Book) {
	b.Item = append(b.Item, input)
}
