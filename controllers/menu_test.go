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
	"github.com/ad3n/resto/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func Test_Menu_Get_Not_Number(t *testing.T) {
	controller := Menu{}

	app := fiber.New()
	app.Get("/menus/:id", controller.Get)

	response, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/menus/abc", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Menu_Get_Menu_Not_Found(t *testing.T) {
	repository := mocks.MenuRepository{}
	repository.On("Find", 123).Return(models.Menu{}, errors.New("triggered")).Once()

	service := services.Menu{Repository: &repository}

	controller := Menu{Service: service}

	app := fiber.New()
	app.Get("/menus/:id", controller.Get)

	response, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/menus/123", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusNotFound, response.StatusCode)
}

func Test_Menu_Get_Success(t *testing.T) {
	repository := mocks.MenuRepository{}
	repository.On("Find", 123).Return(models.Menu{}, nil).Once()

	service := services.Menu{Repository: &repository}

	controller := Menu{Service: service}

	app := fiber.New()
	app.Get("/menus/:id", controller.Get)

	response, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/menus/123", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusOK, response.StatusCode)
}

func Test_Menu_GetAll_Failed(t *testing.T) {
	repository := mocks.MenuRepository{}
	repository.On("FindAll").Return([]models.Menu{}, errors.New("triggered")).Once()

	service := services.Menu{Repository: &repository}

	controller := Menu{Service: service}

	app := fiber.New()
	app.Get("/menus", controller.GetAll)

	response, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/menus", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusInternalServerError, response.StatusCode)
}

func Test_Menu_GetAll_Success(t *testing.T) {
	repository := mocks.MenuRepository{}
	repository.On("FindAll").Return([]models.Menu{}, nil).Once()

	service := services.Menu{Repository: &repository}

	controller := Menu{Service: service}

	app := fiber.New()
	app.Get("/menus", controller.GetAll)

	response, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/menus", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusOK, response.StatusCode)
}

func Test_Menu_Register_Empty_Payload(t *testing.T) {
	controller := Menu{}

	app := fiber.New()
	app.Post("/menus", controller.Create)

	response, err := app.Test(httptest.NewRequest(fiber.MethodPost, "/menus", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Menu_Register_Invalid_Payload(t *testing.T) {
	data := models.Menu{
		Type: types.MENU_DRINK,
		Name: "Susu Jahe",
	}

	repository := mocks.MenuRepository{}

	service := services.Menu{Repository: &repository}

	controller := Menu{Service: service}

	app := fiber.New()
	app.Post("/menus", controller.Create)

	body, err := json.Marshal(data)

	utils.AssertEqual(t, nil, err)

	request := httptest.NewRequest(fiber.MethodPost, "/menus", bytes.NewReader(body))
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusBadRequest, response.StatusCode)
}

func Test_Menu_Register_Valid_Payload_With_Error(t *testing.T) {
	data := models.Menu{
		Type:  types.MENU_DRINK,
		Name:  "Susu Jahe",
		Price: 5000,
	}

	repository := mocks.MenuRepository{}
	repository.On("Saves", &data).Return(errors.New("triggered")).Once()

	service := services.Menu{Repository: &repository}

	controller := Menu{Service: service}

	app := fiber.New()
	app.Post("/menus", controller.Create)

	body, err := json.Marshal(data)

	utils.AssertEqual(t, nil, err)

	request := httptest.NewRequest(fiber.MethodPost, "/menus", bytes.NewReader(body))
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusInternalServerError, response.StatusCode)
}

func Test_Menu_Register_Valid_Payload(t *testing.T) {
	data := models.Menu{
		Type:  types.MENU_DRINK,
		Name:  "Susu Jahe",
		Price: 5000,
	}

	repository := mocks.MenuRepository{}
	repository.On("Saves", &data).Return(nil).Once()

	service := services.Menu{Repository: &repository}

	controller := Menu{Service: service}

	app := fiber.New()
	app.Post("/menus", controller.Create)

	body, err := json.Marshal(data)

	utils.AssertEqual(t, nil, err)

	request := httptest.NewRequest(fiber.MethodPost, "/menus", bytes.NewReader(body))
	request.Header.Add("content-type", "application/json")
	response, err := app.Test(request)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusCreated, response.StatusCode)
}
