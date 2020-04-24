package services

import (
	"application/models"
	"application/repositories"
)

type serviceLayer struct {
	userRepo repositories.IUserRepository
}

// NewUserService returns a new serviceLayer injected with the userRepo
func NewUserService(userRepo repositories.IUserRepository) IServiceLayer {
	return &serviceLayer{userRepo}
}

func (serviceLayer *serviceLayer) CreateUser(userToStore models.User) (models.User, error) {

	user, err := serviceLayer.userRepo.Create(userToStore)
	if err != nil {
		return user, err
	}
	return user, nil
}
