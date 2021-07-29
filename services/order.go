package services

import (
	"errors"
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
	order, err := s.Repository.Find(id)
	if err != nil {
		return order, err
	}

	details, err := s.Detail.FindByOrder(order)
	if err != nil {
		return order, err
	}

	for _, d := range details {
		order.Detail = append(order.Detail, &d)
	}

	return order, err
}

func (s Order) Prepare(order models.Order) (models.Order, error) {
	if order.Status != types.ORDER_PENDING {
		return order, errors.New("can not preparing unreserved order")
	}

	if len(order.Detail) == 0 {
		return order, errors.New("can not preparing empty order")
	}

	order.Status = types.ORDER_PREPARE

	return order, s.Repository.Saves(&order)
}

func (s Order) Served(order models.Order) (models.Order, error) {
	if order.Status != types.ORDER_PREPARE {
		return order, errors.New("can not serving unprepared order")
	}

	order.Status = types.ORDER_SERVED

	return order, s.Repository.Saves(&order)
}

func (s Order) Pay(order models.Order) (models.Order, error) {
	if order.Status != types.ORDER_SERVED {
		return order, errors.New("can not pay unserved order")
	}

	order.Status = types.ORDER_PAID

	return order, s.Repository.Saves(&order)
}

func (s Order) Cancel(order models.Order) (models.Order, error) {
	if order.Status != types.ORDER_PENDING {
		return order, errors.New("can not cancelling unreserved order")
	}

	order.Status = types.ORDER_CANCELLED

	return order, s.Repository.Saves(&order)
}

func (s Order) Rollback(order models.Order) (models.Order, error) {
	if order.Status != types.ORDER_PREPARE {
		return order, errors.New("can not rollback unprepared order")
	}

	order.Status = types.ORDER_PENDING

	return order, s.Repository.Saves(&order)
}

func (s Order) Update(order models.Order) (models.Order, error) {
	if order.Status != types.ORDER_PENDING {
		return order, errors.New("can not update unreserved order")
	}

	order.Amount = 0
	for k, v := range order.Detail {
		menu, err := s.Menu.Find(v.MenuID)
		if err != nil {
			return order, fmt.Errorf("menu with id %d not found", v.MenuID)
		}

		v.OrderID = order.ID

		order.Amount = order.Amount + menu.Price
		order.Detail[k] = v
	}

	err := s.Detail.Saves(order.Detail...)

	return order, err
}
