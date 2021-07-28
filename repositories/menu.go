package repositories

import (
	"errors"

	"github.com/ad3n/resto/models"
	"gorm.io/gorm"
)

type Menu struct {
	Storage *gorm.DB
}

func (r Menu) Find(id int) (models.Menu, error) {
	menu := models.Menu{}
	err := r.Storage.First(&menu, "id = ?", id).Error

	if menu.ID == 0 {
		return menu, errors.New("menu not found")
	}

	return menu, err
}

func (r Menu) FindAll() ([]models.Menu, error) {
	menus := []models.Menu{}
	result := r.Storage.Find(&menus)

	return menus, result.Error
}

func (r Menu) Saves(menus ...*models.Menu) error {
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

	for _, m := range menus {
		err = tx.Save(m).Error
		if err != nil {
			tx.Rollback()

			return err
		}
	}

	return tx.Commit().Error
}
