package repositories

import (
	"errors"

	"github.com/ad3n/resto/models"
	"gorm.io/gorm"
)

type Order struct {
	Storage *gorm.DB
}

func (r Order) Find(id int) (models.Order, error) {
	order := models.Order{}
	err := r.Storage.First(&order, "id = ?", id).Error

	if order.ID == 0 {
		return order, errors.New("order not found")
	}

	return order, err
}

func (r Order) Saves(orders ...*models.Order) error {
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

	for _, m := range orders {
		err = tx.Save(m).Error
		if err != nil {
			tx.Rollback()

			return err
		}
	}

	return tx.Commit().Error
}
