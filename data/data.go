package data

import "github.com/MarioTiara/Go-API-Gin/model"

var DbBooks = model.Books{
	Item: []model.Book{
		{Id: 0, Code: "Code", Title: "Clean Code", Author: "Robert C.Martin", Page: 750, Release: 1990},
		{Id: 1, Code: "ItA", Title: "Introduction to Algorithms", Author: "Thomas H. Cormen", Page: 500, Release: 1993},
		{Id: 2, Code: "TCC", Title: "The Clean Coder", Author: "Robert C.MArtin", Page: 450, Release: 1997},
	},
}
