package book

import "encoding/json"

type BookRequest struct {
	ID          json.Number `json:"id"`
	Title       string      `josn:"title" binding:"required"`
	Description string      `json:"description"`
	Price       json.Number `json:"price" binding:"required"`
	Rating      json.Number `json:"rating" binding:"required"`
}
