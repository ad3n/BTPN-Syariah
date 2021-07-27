package models

type OrderDetail struct {
	ID      int `gorm:"column:id" json:"id"`
	OrderID int `gorm:"column:order_id" json:"order_id"`
	MenuID  int `gorm:"column:menu_id" json:"menu_id" validate:"required"`
}
