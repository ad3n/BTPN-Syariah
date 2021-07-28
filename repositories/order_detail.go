package repositories

import (
	"github.com/ad3n/resto/models"
	"gorm.io/gorm"
)

type OrderDetail struct {
	Storage *gorm.DB
}

func (r OrderDetail) FindByOrder(order models.Order) ([]models.OrderDetail, error) {
	detail := []models.OrderDetail{}
	err := r.Storage.Find(&detail, "order_id = ?", order.ID).Error

	return detail, err
}

func (r OrderDetail) Saves(details ...*models.OrderDetail) error {
	tx := r.Storage.Begin()
	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Error
	if err != nil {
		return err
	}

	for _, m := range details {
		err = tx.Save(m).Error
		if err != nil {
			tx.Rollback()

			return err
		}
	}

	return tx.Commit().Error
}
