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

func (s Menu) Get(id int) (models.Menu, error) {
	return s.Repository.Find(id)
}

func (s Menu) Save(menu models.Menu) (models.Menu, error) {
	return menu, s.Repository.Saves(&menu)
}
