package services

import (
	"errors"
	"testing"

	mocks "github.com/ad3n/resto/mocks/repositories"
	"github.com/ad3n/resto/models"
	"github.com/ad3n/resto/types"
	"github.com/stretchr/testify/assert"
)

func Test_Customer_Not_Found(t *testing.T) {
	repository := mocks.CustomerRepository{}
	repository.On("Find", 1).Return(models.Customer{}, errors.New("customer not found")).Once()

	service := Customer{Repository: &repository}
	_, err := service.Get(1)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Error(), "customer not found")
}

func Test_Customer_Found(t *testing.T) {
	customer := models.Customer{
		ID:          1,
		Name:        "Surya",
		PhoneNumber: "087800093915",
	}

	repository := mocks.CustomerRepository{}
	repository.On("Find", customer.ID).Return(customer, nil).Once()

	service := Customer{Repository: &repository}
	result, err := service.Get(customer.ID)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, customer.Name, result.Name)
}

func Test_Customer_Reservation(t *testing.T) {
	customer := models.Customer{
		ID:          1,
		Name:        "Surya",
		PhoneNumber: "087800093915",
	}

	order := models.Order{
		ID:          0,
		CustomerID:  customer.ID,
		TableNumber: 1,
		Status:      types.ORDER_PENDING,
		Detail:      []*models.OrderDetail{},
	}

	repository := mocks.OrderRepository{}
	repository.On("Saves", &order).Return(nil).Once()

	service := Customer{Order: &repository}
	result, err := service.Reservation(customer, order.TableNumber)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, order.Status, result.Status)
}

func Test_Customer_Save_Phone_Number_Not_Exist(t *testing.T) {
	customer := models.Customer{
		Name:        "Surya",
		PhoneNumber: "087800093915",
	}

	repository := mocks.CustomerRepository{}
	repository.On("Saves", &customer).Return(nil).Once()
	repository.On("FindByPhoneNumber", customer.PhoneNumber).Return(models.Customer{}, nil).Once()

	service := Customer{Repository: &repository}
	_, err := service.Save(customer)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
}

func Test_Customer_Save_Phone_Number_Exist(t *testing.T) {
	customer := models.Customer{
		Name:        "Surya",
		PhoneNumber: "087800093915",
	}

	repository := mocks.CustomerRepository{}
	repository.On("Saves", &customer).Return(nil).Once()
	repository.On("FindByPhoneNumber", customer.PhoneNumber).Return(customer, nil).Once()

	service := Customer{Repository: &repository}
	_, err := service.Save(customer)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
}
