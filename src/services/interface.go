package services

import (
	"application/models"
)

type IServiceLayer interface {
	CreateUser(models.User) (models.User, error)
	// Login()
	// Logout()
	// ChangePassword()
	// IsAmongWinners(uint) (bool, error)
	// PositionInQueue(models.User) (int, error)
}
