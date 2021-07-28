package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	mocks "github.com/ad3n/resto/mocks/repositories"
	"github.com/ad3n/resto/models"
	"github.com/ad3n/resto/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func Test_Customer_Register_Empty_Payload(t *testing.T) {
	repository := mocks.CustomerRepository{}

	service := services.Customer{Repository: &repository}

	controller := Customer{Service: service}

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
