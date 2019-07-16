package model

import "time"

// User model
type User struct {
	ID        uint64     `json:"id" gorm:"column:id"`
	Name      string     `json:"name" gorm:"column:name"`
	SkinType  string     `json:"skinType" gorm:"column:skin_type"`
	Age       uint64     `json:"age" gorm:"column:age"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at"`
}

func (User) TableName() string {
	return "user_users"
}
