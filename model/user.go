package model

import "time"

type User struct {
	ID           int       `json:"_id"`
	FirstName    *string   `json:"fisrt_name"`
	LastName     *string   `json:"last_name"`
	Password     *string   `json:"password" validate:"required, min=2, max=100"`
	Email        *string   `json:"email" validate:"email, required"`
	Phone        *string   `json:"phone" validate:"email, required"`
	Token        *string   `json:"token"`
	UserType     *string   `json:"user_type" validate:"required ,eq=ADMIN|eq=USER"`
	RefreshToken *string   `json:"refresh_token"`
	Created_at   time.Time `json:"created_at"`
	Update_at    time.Time `json:"update_at"`
	User_id      string    `json:"user_id"`
}
