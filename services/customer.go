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

func (s Customer) Get(id int) (models.Customer, error) {
	return s.Repository.Find(id)
}

func (s Customer) Reservation(id int, tableNumber int) (models.Order, error) {
	order := models.Order{
		CustomerID:  id,
		TableNumber: tableNumber,
		Status:      types.ORDER_PENDING,
	}

	return order, s.Order.Saves(&order)
}

func (s Customer) Save(customer models.Customer) (models.Customer, error) {
	exist, _ := s.Repository.FindByPhoneNumber(customer.PhoneNumber)
	if exist.PhoneNumber == customer.PhoneNumber {
		exist.Name = customer.Name
	}

	err := s.Repository.Saves(&customer)

	return customer, err
}
