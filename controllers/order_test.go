package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	mocks "github.com/ad3n/resto/mocks/repositories"
	"github.com/ad3n/resto/models"
	"github.com/ad3n/resto/services"
	"github.com/ad3n/resto/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func Test_Order_Prepare_Id_Not_Number(t *testing.T) {
	controller := Order{}

	app := fiber.New()
	app.Put("/orders/:id/prepare", controller.Prepare)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/abc/prepare", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Order_Prepare_Order_Not_Found(t *testing.T) {
	repository := mocks.OrderRepository{}
	repository.On("Find", 123).Return(models.Order{}, errors.New("triggered")).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id/prepare", controller.Prepare)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/123/prepare", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusNotFound, response.StatusCode)
}

func Test_Order_Prepare_Invalid_Status(t *testing.T) {
	data := models.Order{
		ID:     123,
		Status: types.ORDER_CANCELLED,
	}

	repository := mocks.OrderRepository{}
	repository.On("Find", data.ID).Return(data, nil).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id/prepare", controller.Prepare)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/123/prepare", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Order_Prepare_Success(t *testing.T) {
	data := models.Order{
		ID:     123,
		Status: types.ORDER_PENDING,
	}

	param := models.Order{
		ID:     data.ID,
		Status: types.ORDER_PREPARE,
	}

	repository := mocks.OrderRepository{}
	repository.On("Find", data.ID).Return(data, nil).Once()
	repository.On("Saves", &param).Return(nil).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id/prepare", controller.Prepare)

	request := httptest.NewRequest(fiber.MethodPut, "/orders/123/prepare", nil)
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)

	body, err := ioutil.ReadAll(response.Body)

	utils.AssertEqual(t, nil, err)

	order := models.Order{}
	err = json.Unmarshal(body, &order)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusOK, response.StatusCode)
	utils.AssertEqual(t, param.Status, order.Status)
}

func Test_Order_Cancel_Id_Not_Number(t *testing.T) {
	controller := Order{}

	app := fiber.New()
	app.Put("/orders/:id/cancel", controller.Cancel)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/abc/cancel", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Order_Cancel_Order_Not_Found(t *testing.T) {
	repository := mocks.OrderRepository{}
	repository.On("Find", 123).Return(models.Order{}, errors.New("triggered")).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id/cancel", controller.Cancel)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/123/cancel", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusNotFound, response.StatusCode)
}

func Test_Order_Cancel_Invalid_Status(t *testing.T) {
	data := models.Order{
		ID:     123,
		Status: types.ORDER_CANCELLED,
	}

	repository := mocks.OrderRepository{}
	repository.On("Find", data.ID).Return(data, nil).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id/cancel", controller.Cancel)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/123/cancel", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Order_Cancel_Success(t *testing.T) {
	data := models.Order{
		ID:     123,
		Status: types.ORDER_PENDING,
	}

	param := models.Order{
		ID:     data.ID,
		Status: types.ORDER_CANCELLED,
	}

	repository := mocks.OrderRepository{}
	repository.On("Find", data.ID).Return(data, nil).Once()
	repository.On("Saves", &param).Return(nil).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id/cancel", controller.Cancel)

	request := httptest.NewRequest(fiber.MethodPut, "/orders/123/cancel", nil)
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)

	body, err := ioutil.ReadAll(response.Body)

	utils.AssertEqual(t, nil, err)

	order := models.Order{}
	err = json.Unmarshal(body, &order)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusOK, response.StatusCode)
	utils.AssertEqual(t, param.Status, order.Status)
}

func Test_Order_Rollback_Id_Not_Number(t *testing.T) {
	controller := Order{}

	app := fiber.New()
	app.Put("/orders/:id/rollback", controller.Rollback)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/abc/rollback", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Order_Rollback_Order_Not_Found(t *testing.T) {
	repository := mocks.OrderRepository{}
	repository.On("Find", 123).Return(models.Order{}, errors.New("triggered")).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id/rollback", controller.Rollback)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/123/rollback", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusNotFound, response.StatusCode)
}

func Test_Order_Rollback_Invalid_Status(t *testing.T) {
	data := models.Order{
		ID:     123,
		Status: types.ORDER_CANCELLED,
	}

	repository := mocks.OrderRepository{}
	repository.On("Find", data.ID).Return(data, nil).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id/rollback", controller.Rollback)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/123/rollback", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Order_Rollback_Success(t *testing.T) {
	data := models.Order{
		ID:     123,
		Status: types.ORDER_PREPARE,
	}

	param := models.Order{
		ID:     data.ID,
		Status: types.ORDER_PENDING,
	}

	repository := mocks.OrderRepository{}
	repository.On("Find", data.ID).Return(data, nil).Once()
	repository.On("Saves", &param).Return(nil).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id/rollback", controller.Rollback)

	request := httptest.NewRequest(fiber.MethodPut, "/orders/123/rollback", nil)
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)

	body, err := ioutil.ReadAll(response.Body)

	utils.AssertEqual(t, nil, err)

	order := models.Order{}
	err = json.Unmarshal(body, &order)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusOK, response.StatusCode)
	utils.AssertEqual(t, param.Status, order.Status)
}

func Test_Order_Served_Id_Not_Number(t *testing.T) {
	controller := Order{}

	app := fiber.New()
	app.Put("/orders/:id/served", controller.Served)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/abc/served", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Order_Served_Order_Not_Found(t *testing.T) {
	repository := mocks.OrderRepository{}
	repository.On("Find", 123).Return(models.Order{}, errors.New("triggered")).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id/served", controller.Served)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/123/served", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusNotFound, response.StatusCode)
}

func Test_Order_Served_Invalid_Status(t *testing.T) {
	data := models.Order{
		ID:     123,
		Status: types.ORDER_CANCELLED,
	}

	repository := mocks.OrderRepository{}
	repository.On("Find", data.ID).Return(data, nil).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id/served", controller.Served)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/123/served", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Order_Served_Success(t *testing.T) {
	data := models.Order{
		ID:     123,
		Status: types.ORDER_PREPARE,
	}

	param := models.Order{
		ID:     data.ID,
		Status: types.ORDER_SERVED,
	}

	repository := mocks.OrderRepository{}
	repository.On("Find", data.ID).Return(data, nil).Once()
	repository.On("Saves", &param).Return(nil).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id/served", controller.Served)

	request := httptest.NewRequest(fiber.MethodPut, "/orders/123/served", nil)
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)

	body, err := ioutil.ReadAll(response.Body)

	utils.AssertEqual(t, nil, err)

	order := models.Order{}
	err = json.Unmarshal(body, &order)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusOK, response.StatusCode)
	utils.AssertEqual(t, param.Status, order.Status)
}

func Test_Order_Pay_Id_Not_Number(t *testing.T) {
	controller := Order{}

	app := fiber.New()
	app.Put("/orders/:id/pay", controller.Pay)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/abc/pay", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Order_Pay_Order_Not_Found(t *testing.T) {
	repository := mocks.OrderRepository{}
	repository.On("Find", 123).Return(models.Order{}, errors.New("triggered")).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id/pay", controller.Pay)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/123/pay", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusNotFound, response.StatusCode)
}

func Test_Order_Pay_Invalid_Status(t *testing.T) {
	data := models.Order{
		ID:     123,
		Status: types.ORDER_CANCELLED,
	}

	repository := mocks.OrderRepository{}
	repository.On("Find", data.ID).Return(data, nil).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id/pay", controller.Pay)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/123/pay", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Order_Pay_Success(t *testing.T) {
	data := models.Order{
		ID:     123,
		Status: types.ORDER_SERVED,
	}

	param := models.Order{
		ID:     data.ID,
		Status: types.ORDER_PAID,
	}

	repository := mocks.OrderRepository{}
	repository.On("Find", data.ID).Return(data, nil).Once()
	repository.On("Saves", &param).Return(nil).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id/pay", controller.Pay)

	request := httptest.NewRequest(fiber.MethodPut, "/orders/123/pay", nil)
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)

	body, err := ioutil.ReadAll(response.Body)

	utils.AssertEqual(t, nil, err)

	order := models.Order{}
	err = json.Unmarshal(body, &order)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusOK, response.StatusCode)
	utils.AssertEqual(t, param.Status, order.Status)
}

func Test_Order_Update_Id_Not_Number(t *testing.T) {
	controller := Order{}

	app := fiber.New()
	app.Put("/orders/:id", controller.Update)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/abc", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Order_Update_Order_Not_Found(t *testing.T) {
	repository := mocks.OrderRepository{}
	repository.On("Find", 123).Return(models.Order{}, errors.New("triggered")).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id", controller.Update)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/orders/123", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusNotFound, response.StatusCode)
}

func Test_Order_Update_Order_Invalid_Payload(t *testing.T) {
	data := models.Order{
		ID:     123,
		Status: types.ORDER_CANCELLED,
	}

	repository := mocks.OrderRepository{}
	repository.On("Find", data.ID).Return(data, nil).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id", controller.Update)

	details := []*models.OrderDetail{}

	body, err := json.Marshal(details)

	utils.AssertEqual(t, nil, err)

	request := httptest.NewRequest(fiber.MethodPut, "/orders/123", bytes.NewReader(body))
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Order_Update_Order_Valid_Payload_With_Invalid_Status(t *testing.T) {
	data := models.Order{
		ID:          123,
		TableNumber: 1,
		Status:      types.ORDER_CANCELLED,
	}

	repository := mocks.OrderRepository{}
	repository.On("Find", data.ID).Return(data, nil).Once()

	service := services.Order{Repository: &repository}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id", controller.Update)

	details := []*models.OrderDetail{}
	details = append(details, &models.OrderDetail{MenuID: 1})

	body, err := json.Marshal(details)

	utils.AssertEqual(t, nil, err)

	request := httptest.NewRequest(fiber.MethodPut, "/orders/123", bytes.NewReader(body))
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Order_Update_Order_Valid_Payload_With_Menu_Not_Found(t *testing.T) {
	data := models.Order{
		ID:          123,
		TableNumber: 1,
		Status:      types.ORDER_PENDING,
	}

	repository := mocks.OrderRepository{}
	repository.On("Find", data.ID).Return(data, nil).Once()

	menuRepository := mocks.MenuRepository{}
	menuRepository.On("Find", 1).Return(models.Menu{}, errors.New(""))

	service := services.Order{
		Repository: &repository,
		Menu:       &menuRepository,
	}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id", controller.Update)

	details := []*models.OrderDetail{}
	details = append(details, &models.OrderDetail{MenuID: 1})

	body, err := json.Marshal(details)

	utils.AssertEqual(t, nil, err)

	request := httptest.NewRequest(fiber.MethodPut, "/orders/123", bytes.NewReader(body))
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Order_Update_Order_Valid_Payload_With_Save_Fail(t *testing.T) {
	data := models.Order{
		ID:          123,
		TableNumber: 1,
		Status:      types.ORDER_PENDING,
	}

	detail := &models.OrderDetail{
		MenuID:  1,
		OrderID: data.ID,
	}

	repository := mocks.OrderRepository{}
	repository.On("Find", data.ID).Return(data, nil).Once()

	menuRepository := mocks.MenuRepository{}
	menuRepository.On("Find", 1).Return(models.Menu{}, nil)

	detailRepository := mocks.OrderDetailRepository{}
	detailRepository.On("Saves", detail).Return(errors.New("triggered"))

	service := services.Order{
		Repository: &repository,
		Menu:       &menuRepository,
		Detail:     &detailRepository,
	}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id", controller.Update)

	details := []*models.OrderDetail{}
	details = append(details, detail)

	body, err := json.Marshal(details)

	utils.AssertEqual(t, nil, err)

	request := httptest.NewRequest(fiber.MethodPut, "/orders/123", bytes.NewReader(body))
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Order_Update_Order_Valid_Payload(t *testing.T) {
	data := models.Order{
		ID:          123,
		TableNumber: 1,
		Status:      types.ORDER_PENDING,
	}

	detail := &models.OrderDetail{
		MenuID:  1,
		OrderID: data.ID,
	}

	repository := mocks.OrderRepository{}
	repository.On("Find", data.ID).Return(data, nil).Once()

	menuRepository := mocks.MenuRepository{}
	menuRepository.On("Find", 1).Return(models.Menu{}, nil)

	detailRepository := mocks.OrderDetailRepository{}
	detailRepository.On("Saves", detail).Return(nil)

	service := services.Order{
		Repository: &repository,
		Menu:       &menuRepository,
		Detail:     &detailRepository,
	}

	controller := Order{Service: service}

	app := fiber.New()
	app.Put("/orders/:id", controller.Update)

	details := []*models.OrderDetail{}
	details = append(details, detail)

	body, err := json.Marshal(details)

	utils.AssertEqual(t, nil, err)

	request := httptest.NewRequest(fiber.MethodPut, "/orders/123", bytes.NewReader(body))
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusOK, response.StatusCode)
}
