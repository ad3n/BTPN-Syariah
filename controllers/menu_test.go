package controllers

import (
	"errors"
	"net/http/httptest"
	"testing"

	mocks "github.com/ad3n/resto/mocks/repositories"
	"github.com/ad3n/resto/models"
	"github.com/ad3n/resto/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

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
