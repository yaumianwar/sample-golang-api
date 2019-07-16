package model

import "time"

// Skincare Model
type Skincare struct {
	ID          uint64     `json:"id" gorm:"column:id"`
	UserID      uint64     `json:"userId" gorm:"column:user_id"`
	Brand       string     `json:"brand" gorm:"column:brand"`
	Name        string     `json:"name" gorm:"column:name"`
	ProductType string     `json:"productType" gorm:"column:product_type"`
	SkinConcern string     `json:"skinConcern" gorm:"column:skin_concern"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updated_at"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:created_at"`
}

func (Skincare) TableName() string {
	return "user_skin_cares"
}
