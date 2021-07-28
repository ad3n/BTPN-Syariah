package repositories

import (
	"errors"

	"github.com/ad3n/resto/models"
	"gorm.io/gorm"
)

type Customer struct {
	Storage *gorm.DB
}

func (r Customer) Find(id int) (models.Customer, error) {
	customer := models.Customer{}
	err := r.Storage.First(&customer, "id = ?", id).Error

	if customer.ID == 0 {
		return customer, errors.New("customer not found")
	}

	return customer, err
}

func (r Customer) FindByPhoneNumber(phoneNumber string) (models.Customer, error) {
	customer := models.Customer{}
	err := r.Storage.First(&customer, "phone_number = ?", phoneNumber).Error

	if customer.ID == 0 {
		return customer, errors.New("customer not found")
	}

	return customer, err
}

func (r Customer) Saves(customers ...*models.Customer) error {
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

	for _, m := range customers {
		err = tx.Save(m).Error
		if err != nil {
			tx.Rollback()

			return err
		}
	}

	return tx.Commit().Error
}
