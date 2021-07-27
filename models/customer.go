package models

type Customer struct {
	ID          int    `gorm:"column:id" json:"id"`
	Name        string `gorm:"column:name" json:"name" validate:"required"`
	PhoneNumber string `gorm:"column:phone_number" json:"phone_number" validate:"required"`
}
