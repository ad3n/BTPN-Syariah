package repositories

import "github.com/ad3n/resto/models"

type (
	MenuRepository interface {
		Find(Id int) (models.Menu, error)
		FindAll() ([]models.Menu, error)
		Saves(menus ...*models.Menu) error
	}

	CustomerRepository interface {
		Find(Id int) (models.Customer, error)
		FindAll() ([]models.Customer, error)
		Saves(customers ...*models.Customer) error
	}

	OrderRepository interface {
		Find(Id int) (models.Order, error)
		FindByCustomer(customer models.Customer, status string) ([]models.Order, error)
		FindAll() ([]models.Order, error)
		Saves(orders ...*models.Order) error
	}

	OrderDetailRepository interface {
		FindByOrder(order models.Order) ([]models.OrderDetail, error)
		FindAll() ([]models.OrderDetail, error)
		Saves(details ...*models.OrderDetail) error
	}
)
