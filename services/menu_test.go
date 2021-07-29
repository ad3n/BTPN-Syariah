package services

import (
	"errors"
	"testing"

	mocks "github.com/ad3n/resto/mocks/repositories"
	"github.com/ad3n/resto/models"
	"github.com/ad3n/resto/types"
	"github.com/stretchr/testify/assert"
)

func Test_Menu_Get_All(t *testing.T) {
	repository := mocks.MenuRepository{}
	repository.On("FindAll").Return([]models.Menu{}, nil).Once()

	service := Menu{Repository: &repository}
	result, err := service.GetAll()

	repository.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, 0, len(result))
}

func Test_Menu_Not_Found(t *testing.T) {
	repository := mocks.MenuRepository{}
	repository.On("Find", 1).Return(models.Menu{}, errors.New("menu not found")).Once()

	service := Menu{Repository: &repository}
	_, err := service.Get(1)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Error(), "menu not found")
}

func Test_Menu_Found(t *testing.T) {
	menu := models.Menu{
		ID:    1,
		Type:  types.MENU_MAIN_COURSE,
		Name:  "Ayam Bakar",
		Price: 10000.0,
	}

	repository := mocks.MenuRepository{}
	repository.On("Find", menu.ID).Return(menu, nil).Once()

	service := Menu{Repository: &repository}
	result, err := service.Get(menu.ID)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, menu.Name, result.Name)
}

func Test_Menu_Save(t *testing.T) {
	menu := models.Menu{
		Type:  types.MENU_MAIN_COURSE,
		Name:  "Ayam Bakar",
		Price: 10000.0,
	}

	repository := mocks.MenuRepository{}
	repository.On("Saves", &menu).Return(nil).Once()

	service := Menu{Repository: &repository}
	_, err := service.Save(menu)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
}
