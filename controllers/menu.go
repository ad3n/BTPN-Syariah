package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ad3n/resto/models"
	"github.com/ad3n/resto/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iancoleman/strcase"
)

type Menu struct {
	Service services.Menu
}

// @Description Create new menu (type: drink, main_course, starter, side_dish, snack)
// @Summary Create new menu
// @Tags Menu
// @Accept json
// @Produce json
// @Param menu body forms.MenuCreate true "Menu"
// @Success 200 {object} models.Menu
// @Router /menus [post]
func (c Menu) Create(ctx *fiber.Ctx) error {
	model := models.Menu{}
	ctx.BodyParser(&model)
	model.ID = 0

	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		messages := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, fmt.Sprintf("%s is %s", strcase.ToSnake(err.Field()), err.Tag()))
		}

		ctx.JSON(map[string]interface{}{
			"message": messages,
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	model, err = c.Service.Save(model)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	ctx.JSON(model)

	return ctx.SendStatus(fiber.StatusCreated)
}

// @Description Get menu by ID
// @Summary Get menu by ID
// @Tags Menu
// @Accept json
// @Produce json
// @Param id path int true "Menu ID"
// @Success 200 {object} []models.Menu
// @Router /menus/{id} [get]
func (c Menu) Get(ctx *fiber.Ctx) error {
	menuId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "id is not number",
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	menu, err := c.Service.Get(menuId)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "menu not found",
		})

		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.JSON(menu)
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
