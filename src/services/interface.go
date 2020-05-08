package services

import (
	"application/utils/dto"
)

type IServiceLayer interface {
	CreateUser(dto.UserCreateRequest) (dto.UserCreateResponse, error)
	// Login()
	// Logout()
	// ChangePassword()
	// IsAmongWinners(uint) (bool, error)
	// PositionInQueue(models.User) (int, error)
}
