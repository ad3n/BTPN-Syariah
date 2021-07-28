package services

import (
	"errors"
	"testing"

	mocks "github.com/ad3n/resto/mocks/repositories"
	"github.com/ad3n/resto/models"
	"github.com/ad3n/resto/types"
	"github.com/stretchr/testify/assert"
)

func Test_Order_Not_Found(t *testing.T) {
	repository := mocks.OrderRepository{}
	repository.On("Find", 1).Return(models.Order{}, errors.New("order not found")).Once()

	service := Order{Repository: &repository}
	_, err := service.Get(1)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Error(), "order not found")
}

func Test_Order_Found(t *testing.T) {
	order := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_PENDING,
		Amount:      0,
	}

	repository := mocks.OrderRepository{}
	repository.On("Find", order.ID).Return(order, nil).Once()

	service := Order{Repository: &repository}
	result, err := service.Get(order.ID)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, order.ID, result.ID)
	assert.Equal(t, order.Status, result.Status)
}

func Test_Order_Prepare_Invalid_Status(t *testing.T) {
	order := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_PREPARE,
		Amount:      0,
	}

	service := Order{}
	_, err := service.Prepare(order)

	assert.Equal(t, err.Error(), "can not preparing unreserved order")
}

func Test_Order_Prepare_Valid_Status(t *testing.T) {
	order := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_PENDING,
		Amount:      0,
	}

	param := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_PREPARE,
		Amount:      0,
	}

	repository := mocks.OrderRepository{}
	repository.On("Saves", &param).Return(nil).Once()

	service := Order{Repository: &repository}
	result, err := service.Prepare(order)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, result.Status, param.Status)
}

func Test_Order_Served_Invalid_Status(t *testing.T) {
	order := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_PENDING,
		Amount:      0,
	}

	service := Order{}
	_, err := service.Served(order)

	assert.Equal(t, err.Error(), "can not serving unprepared order")
}

func Test_Order_Served_Valid_Status(t *testing.T) {
	order := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_PREPARE,
		Amount:      0,
	}

	param := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_SERVED,
		Amount:      0,
	}

	repository := mocks.OrderRepository{}
	repository.On("Saves", &param).Return(nil).Once()

	service := Order{Repository: &repository}
	result, err := service.Served(order)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, result.Status, param.Status)
}

func Test_Order_Pay_Invalid_Status(t *testing.T) {
	order := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_PENDING,
		Amount:      0,
	}

	service := Order{}
	_, err := service.Pay(order)

	assert.Equal(t, err.Error(), "can not pay unserved order")
}

func Test_Order_Pay_Valid_Status(t *testing.T) {
	order := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_SERVED,
		Amount:      0,
	}

	param := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_PAID,
		Amount:      0,
	}

	repository := mocks.OrderRepository{}
	repository.On("Saves", &param).Return(nil).Once()

	service := Order{Repository: &repository}
	result, err := service.Pay(order)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, result.Status, param.Status)
}

func Test_Order_Rollback_Invalid_Status(t *testing.T) {
	order := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_PENDING,
		Amount:      0,
	}

	service := Order{}
	_, err := service.Rollback(order)

	assert.Equal(t, err.Error(), "can not rollback unprepared order")
}

func Test_Order_Rollback_Valid_Status(t *testing.T) {
	order := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_PREPARE,
		Amount:      0,
	}

	param := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_PENDING,
		Amount:      0,
	}

	repository := mocks.OrderRepository{}
	repository.On("Saves", &param).Return(nil).Once()

	service := Order{Repository: &repository}
	result, err := service.Rollback(order)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, result.Status, param.Status)
}

func Test_Order_Cancel_Invalid_Status(t *testing.T) {
	order := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_SERVED,
		Amount:      0,
	}

	service := Order{}
	_, err := service.Cancel(order)

	assert.Equal(t, err.Error(), "can not cancelling unreserved order")
}

func Test_Order_Cancel_Valid_Status(t *testing.T) {
	order := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_PENDING,
		Amount:      0,
	}

	param := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_CANCELLED,
		Amount:      0,
	}

	repository := mocks.OrderRepository{}
	repository.On("Saves", &param).Return(nil).Once()

	service := Order{Repository: &repository}
	result, err := service.Cancel(order)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, result.Status, param.Status)
}

func Test_Order_Update_Invalid_Status(t *testing.T) {
	order := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_PREPARE,
		Amount:      0,
	}

	service := Order{}
	_, err := service.Update(order)

	assert.Equal(t, err.Error(), "can not update unreserved order")
}

func Test_Order_Update_Menu_Not_Found(t *testing.T) {
	order := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_PENDING,
		Amount:      0,
	}

	order.Detail = append(order.Detail, &models.OrderDetail{OrderID: 1, MenuID: 1})

	repository := mocks.MenuRepository{}
	repository.On("Find", 1).Return(models.Menu{}, errors.New(""))

	service := Order{Menu: &repository}
	_, err := service.Update(order)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Error(), "menu with id 1 not found")
}

func Test_Order_Update_Valid(t *testing.T) {
	order := models.Order{
		ID:          1,
		CustomerID:  1,
		TableNumber: 1,
		Status:      types.ORDER_PENDING,
		Amount:      0,
	}

	detail := models.OrderDetail{OrderID: 1, MenuID: 1}

	order.Detail = append(order.Detail, &detail)

	menuRepository := mocks.MenuRepository{}
	menuRepository.On("Find", 1).Return(models.Menu{}, nil)

	detailRepository := mocks.OrderDetailRepository{}
	detailRepository.On("Saves", &detail).Return(nil).Once()

	service := Order{Menu: &menuRepository, Detail: &detailRepository}
	_, err := service.Update(order)

	menuRepository.AssertExpectations(t)
	detailRepository.AssertExpectations(t)

	assert.Nil(t, err)
}
