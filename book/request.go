package book

import "encoding/json"

type BookRequest struct {
	Title       string      `josn:"title" binding:"required"`
	Description string      `json:"description"`
	Price       json.Number `json:"price" binding:"required, number"`
	Rating      json.Number `json:"rating" binding:"required"`
}
