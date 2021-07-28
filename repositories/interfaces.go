package repositories

import "github.com/ad3n/resto/models"

type (
	MenuRepository interface {
		Find(id int) (models.Menu, error)
		FindAll() ([]models.Menu, error)
		Saves(menus ...*models.Menu) error
	}

	CustomerRepository interface {
		Find(id int) (models.Customer, error)
		FindByPhoneNumber(phoneNumber string) (models.Customer, error)
		Saves(customers ...*models.Customer) error
	}

	OrderRepository interface {
		Find(id int) (models.Order, error)
		Saves(orders ...*models.Order) error
	}

	OrderDetailRepository interface {
		FindByOrder(order models.Order) ([]models.OrderDetail, error)
		Saves(details ...*models.OrderDetail) error
	}
)
