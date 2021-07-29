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

type Order struct {
	Service services.Order
}

// @Description Send order to kitchen (only for reservation - pending - order)
// @Summary Send order to kitchen
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Router /orders/{id}/prepare [put]
func (c Order) Prepare(ctx *fiber.Ctx) error {
	return c.handleStatus(ctx, "prepare")
}

// @Description Cancel reservation (only for reservation - pending - order)
// @Summary Cancel reservation
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Router /orders/{id}/cancel [put]
func (c Order) Cancel(ctx *fiber.Ctx) error {
	return c.handleStatus(ctx, "cancel")
}

// @Description Send back to waiter because can't be created (only for prapare - prepare - order)
// @Summary Send back to waiter
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Router /orders/{id}/rollback [put]
func (c Order) Rollback(ctx *fiber.Ctx) error {
	return c.handleStatus(ctx, "rollback")
}

// @Description Serving order (only for prapare - prepare - order)
// @Summary Serving order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Router /orders/{id}/served [put]
func (c Order) Served(ctx *fiber.Ctx) error {
	return c.handleStatus(ctx, "served")
}

// @Description Order payment - completed - (only for served - served - order)
// @Summary Order payment
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Router /orders/{id}/pay [put]
func (c Order) Pay(ctx *fiber.Ctx) error {
	return c.handleStatus(ctx, "pay")
}

// @Description Get order by ID
// @Summary Get order by ID
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} []models.Order
// @Router /orders/{id} [get]
func (c Order) Get(ctx *fiber.Ctx) error {
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
			"message": "order not found",
		})

		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.JSON(order)
}

// @Description Update order (assign/update menu to order)
// @Summary Update order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Router /orders/{id} [put]
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

		return ctx.SendStatus(http.StatusNotFound)
	}

	model := []*models.OrderDetail{}
	ctx.BodyParser(&model)

	order.Detail = model

	validate := validator.New()
	err = validate.Struct(order)
	if err != nil {
		messages := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, fmt.Sprintf("%s is %s", strcase.ToSnake(err.Field()), err.Tag()))
		}

		fmt.Println(messages)

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

func (c Order) handleStatus(ctx *fiber.Ctx, status string) error {
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

		return ctx.SendStatus(http.StatusNotFound)
	}

	switch status {
	case "prepare":
		order, err = c.Service.Prepare(order)
	case "cancel":
		order, err = c.Service.Cancel(order)
	case "rollback":
		order, err = c.Service.Rollback(order)
	case "served":
		order, err = c.Service.Served(order)
	case "pay":
		order, err = c.Service.Pay(order)
	}

	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(order)
}
