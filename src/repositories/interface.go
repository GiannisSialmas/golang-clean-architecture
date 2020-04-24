package repositories

import (
	"application/models"
)

// The interface that all repositories should conform to
type IUserRepository interface {
	Create(user models.User) (models.User, error)
	GetOne(id uint) (models.User, error)
	GetAll() ([]models.User, error)
	GetByEmail(email string) (models.User, error)
	// Update(user models.User) (models.User, error)
	Delete(id uint) error
}
