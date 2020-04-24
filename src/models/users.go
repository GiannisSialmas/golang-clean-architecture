package models

import (
	"time"
)

// User represents an object as mapped from a row in the database Users table
type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstName" gorm:"type:varchar"`
	LastName  string `json:"lastName" gorm:"type:varchar"`
	Email     string `json:"email" gorm:"type:varchar;unique_index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
