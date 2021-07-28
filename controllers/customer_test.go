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

func Test_Customer_Register_Empty_Payload(t *testing.T) {
	controller := Customer{}

	app := fiber.New()
	app.Post("/customers", controller.Register)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPost, "/customers", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Customer_Register_Invalid_Payload(t *testing.T) {
	data := models.Customer{
		Name: "Surya",
	}

	repository := mocks.CustomerRepository{}

	service := services.Customer{Repository: &repository}

	controller := Customer{Service: service}

	app := fiber.New()
	app.Post("/customers", controller.Register)

	body, err := json.Marshal(data)

	utils.AssertEqual(t, nil, err)

	request := httptest.NewRequest(fiber.MethodPost, "/customers", bytes.NewReader(body))
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Customer_Register_Valid_Payload_With_Error(t *testing.T) {
	data := models.Customer{
		Name:        "Surya",
		PhoneNumber: "087800093915",
	}

	repository := mocks.CustomerRepository{}
	repository.On("Saves", &data).Return(errors.New("triggered")).Once()
	repository.On("FindByPhoneNumber", data.PhoneNumber).Return(data, nil).Once()

	service := services.Customer{Repository: &repository}

	controller := Customer{Service: service}

	app := fiber.New()
	app.Post("/customers", controller.Register)

	body, err := json.Marshal(data)

	utils.AssertEqual(t, nil, err)

	request := httptest.NewRequest(fiber.MethodPost, "/customers", bytes.NewReader(body))
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusInternalServerError, response.StatusCode)
}

func Test_Customer_Register_Valid_Payload(t *testing.T) {
	data := models.Customer{
		Name:        "Surya",
		PhoneNumber: "087800093915",
	}

	repository := mocks.CustomerRepository{}
	repository.On("Saves", &data).Return(nil).Once()
	repository.On("FindByPhoneNumber", data.PhoneNumber).Return(data, nil).Once()

	service := services.Customer{Repository: &repository}

	controller := Customer{Service: service}

	app := fiber.New()
	app.Post("/customers", controller.Register)

	body, err := json.Marshal(data)

	utils.AssertEqual(t, nil, err)

	request := httptest.NewRequest(fiber.MethodPost, "/customers", bytes.NewReader(body))
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusCreated, response.StatusCode)
}

func Test_Customer_Reservation_Id_Not_Number(t *testing.T) {
	controller := Customer{}

	app := fiber.New()
	app.Post("/customers/:id/reservation", controller.Reservation)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPost, "/customers/abc/reservation", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Customer_Reservation_Customer_Not_Found(t *testing.T) {
	repository := mocks.CustomerRepository{}
	repository.On("Find", 123).Return(models.Customer{}, errors.New("triggered")).Once()

	service := services.Customer{Repository: &repository}

	controller := Customer{Service: service}

	app := fiber.New()
	app.Post("/customers/:id/reservation", controller.Reservation)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPost, "/customers/123/reservation", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusNotFound, response.StatusCode)
}

func Test_Customer_Reservation_Empty_Payload(t *testing.T) {
	customer := models.Customer{
		ID:          123,
		Name:        "Surya",
		PhoneNumber: "087800093915",
	}

	repository := mocks.CustomerRepository{}
	repository.On("Find", customer.ID).Return(customer, nil).Once()

	service := services.Customer{Repository: &repository}

	controller := Customer{Service: service}

	app := fiber.New()
	app.Post("/customers/:id/reservation", controller.Reservation)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPost, "/customers/123/reservation", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Customer_Reservation_Invalid_Payload(t *testing.T) {
	data := models.Order{}

	customer := models.Customer{
		ID:          123,
		Name:        "Surya",
		PhoneNumber: "087800093915",
	}

	repository := mocks.CustomerRepository{}
	repository.On("Find", customer.ID).Return(customer, nil).Once()

	service := services.Customer{Repository: &repository}

	controller := Customer{Service: service}

	app := fiber.New()
	app.Post("/customers/:id/reservation", controller.Reservation)

	body, err := json.Marshal(data)

	utils.AssertEqual(t, nil, err)

	request := httptest.NewRequest(fiber.MethodPost, "/customers/123/reservation", bytes.NewReader(body))
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Customer_Reservation_Valid_Payload_With_Error(t *testing.T) {
	customer := models.Customer{
		ID:          123,
		Name:        "Surya",
		PhoneNumber: "087800093915",
	}

	data := models.Order{
		TableNumber: 123,
		CustomerID:  customer.ID,
		Status:      types.ORDER_PENDING,
	}

	repository := mocks.CustomerRepository{}
	repository.On("Find", customer.ID).Return(customer, nil).Once()

	orderRepository := mocks.OrderRepository{}
	orderRepository.On("Saves", &data).Return(errors.New("triggered")).Once()

	service := services.Customer{
		Repository: &repository,
		Order:      &orderRepository,
	}

	controller := Customer{Service: service}

	app := fiber.New()
	app.Post("/customers/:id/reservation", controller.Reservation)

	body, err := json.Marshal(data)

	utils.AssertEqual(t, nil, err)

	request := httptest.NewRequest(fiber.MethodPost, "/customers/123/reservation", bytes.NewReader(body))
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusInternalServerError, response.StatusCode)
}

func Test_Customer_Reservation_Valid_Payload(t *testing.T) {
	customer := models.Customer{
		ID:          123,
		Name:        "Surya",
		PhoneNumber: "087800093915",
	}

	data := models.Order{
		TableNumber: 123,
		CustomerID:  customer.ID,
		Status:      types.ORDER_PENDING,
	}

	repository := mocks.CustomerRepository{}
	repository.On("Find", customer.ID).Return(customer, nil).Once()

	orderRepository := mocks.OrderRepository{}
	orderRepository.On("Saves", &data).Return(nil).Once()

	service := services.Customer{
		Repository: &repository,
		Order:      &orderRepository,
	}

	controller := Customer{Service: service}

	app := fiber.New()
	app.Post("/customers/:id/reservation", controller.Reservation)

	body, err := json.Marshal(data)

	utils.AssertEqual(t, nil, err)

	request := httptest.NewRequest(fiber.MethodPost, "/customers/123/reservation", bytes.NewReader(body))
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)

	body, err = ioutil.ReadAll(response.Body)

	utils.AssertEqual(t, nil, err)

	order := models.Order{}
	err = json.Unmarshal(body, &order)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusCreated, response.StatusCode)
	utils.AssertEqual(t, order.TableNumber, data.TableNumber)
}
