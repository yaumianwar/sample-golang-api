package user

import "github.com/jinzhu/gorm"

// Service for User Package
type Service struct {
	db     *gorm.DB
}

// NewService create new User Service
func NewService(db *gorm.DB) Service {
	return Service{db}
}

