package routes

import (
	"github.com/ad3n/resto/configs"
	"github.com/ad3n/resto/controllers"
	"github.com/ad3n/resto/repositories"
	"github.com/ad3n/resto/services"
	"github.com/gofiber/fiber/v2"
)

type Customer struct {
}

func (Customer) RegisterRoutes(router fiber.Router) {
	customerRepository := repositories.Customer{Storage: configs.Db}
	orderRepository := repositories.Order{Storage: configs.Db}

	customerService := services.Customer{
		Repository: &customerRepository,
		Order:      &orderRepository,
	}

	customer := controllers.Customer{Service: customerService}

	router.Post("/customers", customer.Register)
	router.Post("/customers/:id/reservation", customer.Reservation)
}
