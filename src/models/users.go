package models

import (
	"time"
)

// User represents an object as mapped from a row in the database Users table
type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"type:varchar"`
	LastName  string `gorm:"type:varchar"`
	Email     string `gorm:"type:varchar;unique_index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
