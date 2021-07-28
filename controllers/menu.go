package controllers

import (
	"github.com/ad3n/resto/services"
	"github.com/gofiber/fiber/v2"
)

type Menu struct {
	Service services.Menu
}

func (c Menu) GetAll(ctx *fiber.Ctx) error {
	menus, err := c.Service.GetAll()
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "unable to get all menus",
		})

		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(menus)
}
