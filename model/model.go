package model

type Book struct {
	Id      int
	Title   string
	Author  string
	Page    int
	Release int
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
