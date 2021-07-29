package routes

import (
	"github.com/ad3n/resto/configs"
	"github.com/ad3n/resto/controllers"
	"github.com/ad3n/resto/repositories"
	"github.com/ad3n/resto/services"
	"github.com/gofiber/fiber/v2"
)

type Order struct {
}

func (Order) RegisterRoutes(router fiber.Router) {
	orderRepository := repositories.Order{Storage: configs.Db}
	detailRepostiory := repositories.OrderDetail{Storage: configs.Db}
	menuRepository := repositories.Menu{Storage: configs.Db}

	orderService := services.Order{
		Repository: &orderRepository,
		Detail:     &detailRepostiory,
		Menu:       &menuRepository,
	}

	order := controllers.Order{Service: orderService}

	router.Put("/orders/:id", order.Update)
	router.Get("/orders/:id", order.Get)
	router.Put("/orders/:id/prepare", order.Prepare)
	router.Put("/orders/:id/cancel", order.Cancel)
	router.Put("/orders/:id/rollback", order.Rollback)
	router.Put("/orders/:id/served", order.Served)
	router.Put("/orders/:id/pay", order.Pay)
}
