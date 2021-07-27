package services

import (
	"github.com/ad3n/resto/models"
	"github.com/ad3n/resto/repositories"
	"github.com/ad3n/resto/types"
)

type Customer struct {
	Repository  repositories.CustomerRepository
	Order       repositories.OrderRepository
	OrderDetail repositories.OrderDetailRepository
}

func (s Customer) GetOrder(customer models.Customer, status string) ([]models.Order, error) {
	orders, err := s.Order.FindByCustomer(customer, status)
	if err != nil {
		return orders, err
	}

	for k, o := range orders {
		o.Detail, err = s.OrderDetail.FindByOrder(o)
		if err == nil {
			orders[k] = o
		}
	}

	return orders, err
}

func (s Customer) Get(customerId int) (models.Customer, error) {
	return s.Repository.Find(customerId)
}

func (s Customer) Reservation(customerId int, tableNumber int) (models.Order, error) {
	order := models.Order{
		CustomerID:  customerId,
		TableNumber: tableNumber,
		Status:      types.ORDER_PENDING,
	}

	return order, s.Order.Saves(&order)
}

func (s Customer) Save(customer models.Customer) (models.Customer, error) {
	err := s.Repository.Saves(&customer)

	return customer, err
}
