package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ad3n/resto/models"
	"github.com/ad3n/resto/services"
	"github.com/ad3n/resto/types"
	"github.com/ad3n/resto/utils"
	"github.com/gofiber/fiber/v2"

	"github.com/go-playground/validator/v10"
)

type Customer struct {
	Service services.Customer
}

func (c Customer) Register(ctx *fiber.Ctx) error {
	model := models.Customer{}
	ctx.BodyParser(&model)

	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		messages := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, fmt.Sprintf("%s is %s", utils.ToUnderScore(err.Field()), err.Tag()))
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

func (c Customer) GetOrder(ctx *fiber.Ctx) error {
	customerId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "id is not number",
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	customer, err := c.Service.Get(customerId)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	order, err := c.Service.GetOrder(customer, ctx.Query("status", types.ORDER_SERVED))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(order)
}

func (c Customer) Reservation(ctx *fiber.Ctx) error {
	customerId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "id is not number",
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	customer, err := c.Service.Get(customerId)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	model := models.Order{}
	ctx.BodyParser(&model)

	validate := validator.New()
	err = validate.Struct(model)
	if err != nil {
		messages := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, fmt.Sprintf("%s is %s", utils.ToUnderScore(err.Field()), err.Tag()))
		}

		ctx.JSON(map[string]interface{}{
			"message": messages,
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	model, err = c.Service.Reservation(customer.ID, model.TableNumber)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	ctx.JSON(model)

	return ctx.SendStatus(fiber.StatusCreated)
}
