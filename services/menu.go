package services

import (
	"github.com/ad3n/resto/models"
	"github.com/ad3n/resto/repositories"
)

type Menu struct {
	Repository repositories.MenuRepository
}

func (s Menu) GetAll() ([]models.Menu, error) {
	return s.Repository.FindAll()
}

func (s Menu) Save(menu models.Menu) error {
	return s.Repository.Saves(&menu)
}
