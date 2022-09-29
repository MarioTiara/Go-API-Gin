package user

import "gorm.io/gorm"

type Repository interface {
	FindByName(username string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByName(username string) (User, error) {
	var user User

	err := r.db.Where("username=?", username).First(&user).Error
	return user, err
}
