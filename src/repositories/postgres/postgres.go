package postgres

import (
	"application/exceptions"
	"application/models"
	"application/repositories"
	"application/utils/dto"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository returns a postgres user repo
func NewUserRepository(db *gorm.DB) repositories.IUserRepository {
	return &userRepository{db}
}

func (userRepository *userRepository) Create(user dto.UserCreateRequest) (dto.UserCreateResponse, error) {

	// Check if user Already exists
	if err := userRepository.db.Where("email = ?", user.Email).First(&models.User{}).Error; !gorm.IsRecordNotFoundError(err) {
		return dto.UserCreateResponse{}, exceptions.ErrUserEmailExists
	}

	// DTO to db model for gorm usage
	var userCreated models.User
	copier.Copy(&userCreated, &user)

	if err := userRepository.db.Create(&userCreated).Error; err != nil {
		return dto.UserCreateResponse{}, err
	}

	// Db model to dto to return to upper layers
	var userCreatedResponseDto dto.UserCreateResponse
	copier.Copy(&userCreated, &userCreatedResponseDto)

	// Gorm after creating the record, inserts the ID plus timestamps into the object
	return userCreatedResponseDto, nil

}

func (userRepository *userRepository) GetOne(id uint) (models.User, error) {

	var user models.User
	result := userRepository.db.Where("ID = ?", id).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	// Gorm after creating the record, inserts the ID plus timestamps into the object
	return user, nil

}

func (userRepository *userRepository) GetByEmail(email string) (models.User, error) {

	// TODO: Should i validate the email here?

	var user models.User
	result := userRepository.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, exceptions.ErrGeneral
	} else if result.RecordNotFound() {
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

// func (userRepository *userRepository) Update(user models.User) (models.User, error) {

// 	return user, nil

// }

func (userRepository *userRepository) Delete(id uint) error {

	result := userRepository.db.Where("ID = ?", id).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	// Gorm does not return an error if zero entries are deleted so we must manually throw an error
	if result.RowsAffected == 0 {
		return exceptions.ErrUserNotFound
	}

	return nil

}
