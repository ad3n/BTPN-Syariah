package models

type Order struct {
	ID          int           `gorm:"column:id" json:"id"`
	CustomerID  int           `gorm:"column:customer_id" json:"customer_id"`
	TableNumber int           `gorm:"column:table_number" json:"table_number" validate:"required"`
	Status      string        `gorm:"column:status" json:"status"`
	Amount      float32       `gorm:"column:amount;type:decimal(17,2)" json:"amount"`
	Detail      []OrderDetail `json:"details"`
}
