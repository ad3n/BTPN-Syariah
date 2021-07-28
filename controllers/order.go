package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ad3n/resto/models"
	"github.com/ad3n/resto/services"
	"github.com/ad3n/resto/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Order struct {
	Service services.Order
}

func (c Order) Prepare(ctx *fiber.Ctx) error {
	orderId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "id is not number",
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	order, err := c.Service.Get(orderId)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	order, err = c.Service.Prepare(order)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(order)
}

func (c Order) Cancel(ctx *fiber.Ctx) error {
	orderId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "id is not number",
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	order, err := c.Service.Get(orderId)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	order, err = c.Service.Cancel(order)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(order)
}

func (c Order) Rollback(ctx *fiber.Ctx) error {
	orderId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "id is not number",
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	order, err := c.Service.Get(orderId)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	order, err = c.Service.Rollback(order)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(order)
}

func (c Order) Served(ctx *fiber.Ctx) error {
	orderId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "id is not number",
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	order, err := c.Service.Get(orderId)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	order, err = c.Service.Served(order)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(order)
}

func (c Order) Pay(ctx *fiber.Ctx) error {
	orderId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "id is not number",
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	order, err := c.Service.Get(orderId)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	order, err = c.Service.Pay(order)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(order)
}

func (c Order) Update(ctx *fiber.Ctx) error {
	orderId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "id is not number",
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	order, err := c.Service.Get(orderId)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(http.StatusBadRequest)
	}

	model := []*models.OrderDetail{}
	ctx.BodyParser(&model)

	order.Detail = model

	validate := validator.New()
	err = validate.Struct(order)
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

	order, err = c.Service.Update(order)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(order)
}
