package models

import "github.com/jinzhu/gorm"

// User represents user entities in the system
type User struct {
	gorm.Model
	Name  string
	Email string
}
