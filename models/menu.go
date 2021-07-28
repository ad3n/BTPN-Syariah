package models

type Menu struct {
	ID    int     `gorm:"column:id" json:"id"`
	Type  string  `gorm:"column:type" json:"type"`
	Name  string  `gorm:"column:name" json:"name"`
	Price float32 `gorm:"column:price;type:decimal(17,2)" json:"price"`
}
