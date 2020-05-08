package services

import (
	"application/repositories"
	"application/utils/dto"
)

type serviceLayer struct {
	userRepo repositories.IUserRepository
}

// NewUserService returns a new serviceLayer injected with the userRepo
func NewUserService(userRepo repositories.IUserRepository) IServiceLayer {
	return &serviceLayer{userRepo}
}

func (serviceLayer *serviceLayer) CreateUser(userToStore dto.UserCreateRequest) (dto.UserCreateResponse, error) {

	user, err := serviceLayer.userRepo.Create(userToStore)
	if err != nil {
		return user, err
	}
	return user, nil
}
