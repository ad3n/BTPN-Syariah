package routes

import (
	"github.com/ad3n/resto/configs"
	"github.com/ad3n/resto/controllers"
	"github.com/ad3n/resto/repositories"
	"github.com/ad3n/resto/services"
	"github.com/gofiber/fiber/v2"
)

type Menu struct {
}

func (Menu) RegisterRoutes(router fiber.Router) {
	menuRepository := repositories.Menu{Storage: configs.Db}
	menuService := services.Menu{Repository: &menuRepository}

	menu := controllers.Menu{Service: menuService}

	router.Get("/menus", menu.GetAll)
}
