package postgres

import (
	"application/models"
	"application/repositories"
	"application/repositories/exceptions"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository returns a postgres user repo
func NewUserRepository(db *gorm.DB) repositories.IUserRepository {
	return &userRepository{db}
}

func (userRepository *userRepository) Create(user models.User) (models.User, error) {

	if err := userRepository.db.Create(&user).Error; err != nil {
		return user, err
	}
	// Gorm after creating the record, inserts the ID plus timestamps into the object
	return user, nil

}

func (userRepository *userRepository) GetOne(id uint) (models.User, error) {

	var user models.User
	if err := userRepository.db.First(&user, id).Error; err != nil {
		return user, exceptions.ErrUserNotFound
	}
	// Gorm after creating the record, inserts the ID plus timestamps into the object
	return user, nil

}

func (userRepository *userRepository) GetAll() ([]models.User, error) {

	var users []models.User
	if err := userRepository.db.Find(&users).Error; err != nil {
		return users, err
	}
	// Gorm after creating the record, inserts the ID plus timestamps into the object
	return users, nil

}

func (userRepository *userRepository) Update(user models.User) (models.User, error) {

	return user, nil

}

func (userRepository *userRepository) Delete(id uint) error {

	result := userRepository.db.Where("ID = ?", id).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return exceptions.ErrUserNotFound
	}
	// Gorm after creating the record, inserts the ID plus timestamps into the object
	return nil

}
