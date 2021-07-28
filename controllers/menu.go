package controllers

import (
	"github.com/ad3n/resto/services"
	"github.com/gofiber/fiber/v2"
)

type Menu struct {
	Service services.Menu
}

// @Description Get all menu
// @Summary Get all menu
// @Tags Menu
// @Accept json
// @Produce json
// @Success 200 {object} []models.Menu
// @Router /menus [get]
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
