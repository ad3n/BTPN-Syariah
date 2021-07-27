package models

type OrderDetail struct {
	ID      int    `gorm:"column:id" json:"id"`
	OrderID int    `gorm:"column:order_id" json:"order_id"`
	MenuId  string `gorm:"column:menu_id" json:"menu_id"`
}
