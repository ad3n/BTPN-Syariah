package services

import (
	"fmt"

	"github.com/ad3n/resto/models"
	"github.com/ad3n/resto/repositories"
	"github.com/ad3n/resto/types"
)

type Order struct {
	Repository repositories.OrderRepository
	Detail     repositories.OrderDetailRepository
	Menu       repositories.MenuRepository
}

func (s Order) Get(id int) (models.Order, error) {
	return s.Repository.Find(id)
}

func (s Order) Prepare(order models.Order) (models.Order, error) {
	order.Status = types.ORDER_PREPARE

	return order, s.Repository.Saves(&order)
}

func (s Order) Update(order models.Order) (models.Order, error) {
	for k, v := range order.Detail {
		_, err := s.Menu.Find(v.MenuID)
		if err != nil {
			return order, fmt.Errorf("menu with id %d not found", v.MenuID)
		}

		v.OrderID = order.ID

		order.Detail[k] = v
	}

	err := s.Detail.Saves(order.Detail...)

	return order, err
}
