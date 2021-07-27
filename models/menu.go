package models

type Menu struct {
	ID   int    `gorm:"column:id" json:"id"`
	Type string `gorm:"column:type" json:"type"`
	Name string `gorm:"column:name" json:"name"`
}
