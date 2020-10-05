package service_test

import (
	"testing"

	"github.com/auliman-a/golang-test/graph/model"
	"github.com/auliman-a/golang-test/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockInventoryService struct {
	mock.Mock
}

func (mock *MockInventoryService) Checkout() (*model.Cart, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*model.Cart), args.Error(1)
}
func (mock *MockInventoryService) GetInventoryData() []*model.Item {
	args := mock.Called()
	result := args.Get(0)
	return result.([]*model.Item)
}
func (mock *MockInventoryService) GetCartItem() *model.Cart {
	args := mock.Called()
	result := args.Get(0)
	return result.(*model.Cart)
}
func (mock *MockInventoryService) AddItemToCart(input model.CartInput) error {
	args := mock.Called(input)
	// result := args.Get(0)
	return args.Error(0)
}

func TestService(t *testing.T) {
	t.Run("Test CheckOut with Promo 1", func(t *testing.T) {

		item1 := model.Item{
			Sku:      "120P90",
			Name:     "Google Home",
			Quantity: 10,
			Price:    49.99,
		}

		item2 := model.Item{
			Sku:      "43N23P",
			Name:     "MacBook Pro",
			Quantity: 5,
			Price:    5399.99,
		}

		item3 := model.Item{
			Sku:      "A304SD",
			Name:     "Alexa Speaker",
			Quantity: 10,
			Price:    109.50,
		}

		item4 := model.Item{
			Sku:      "234234",
			Name:     "Raspberry Pi B",
			Quantity: 2,
			Price:    30.00,
		}

		InventoryData := map[string]model.Item{
			item1.Sku: item1,
			item2.Sku: item2,
			item3.Sku: item3,
			item4.Sku: item4,
		}

		cartItem := model.CartOutput{
			Sku:      "43N23P",
			Name:     "",
			Quantity: 1,
			Amount:   5399.99,
		}

		var CartData = model.Cart{}
		CartData.CartItems = append(CartData.CartItems, &cartItem)

		inventoryService := service.NewInventoryService(InventoryData, CartData)

		cartResult, err := inventoryService.Checkout()

		assert.Nil(t, err)
		assert.Equal(t, 5399.99, cartResult.TotalPrice)
	})

	t.Run("Test CheckOut with Promo 2", func(t *testing.T) {

		item1 := model.Item{
			Sku:      "120P90",
			Name:     "Google Home",
			Quantity: 10,
			Price:    49.99,
		}

		item2 := model.Item{
			Sku:      "43N23P",
			Name:     "MacBook Pro",
			Quantity: 5,
			Price:    5399.99,
		}

		item3 := model.Item{
			Sku:      "A304SD",
			Name:     "Alexa Speaker",
			Quantity: 10,
			Price:    109.50,
		}

		item4 := model.Item{
			Sku:      "234234",
			Name:     "Raspberry Pi B",
			Quantity: 2,
			Price:    30.00,
		}

		InventoryData := map[string]model.Item{
			item1.Sku: item1,
			item2.Sku: item2,
			item3.Sku: item3,
			item4.Sku: item4,
		}

		cartItem := model.CartOutput{
			Sku:      "120P90",
			Name:     "",
			Quantity: 3,
			Amount:   149.97,
		}

		var CartData = model.Cart{}
		CartData.CartItems = append(CartData.CartItems, &cartItem)

		inventoryService := service.NewInventoryService(InventoryData, CartData)

		cartResult, err := inventoryService.Checkout()

		assert.Nil(t, err)
		assert.Equal(t, 99.98, cartResult.TotalPrice)
	})

	t.Run("Test CheckOut with Promo 3", func(t *testing.T) {

		item1 := model.Item{
			Sku:      "120P90",
			Name:     "Google Home",
			Quantity: 10,
			Price:    49.99,
		}

		item2 := model.Item{
			Sku:      "43N23P",
			Name:     "MacBook Pro",
			Quantity: 5,
			Price:    5399.99,
		}

		item3 := model.Item{
			Sku:      "A304SD",
			Name:     "Alexa Speaker",
			Quantity: 10,
			Price:    109.50,
		}

		item4 := model.Item{
			Sku:      "234234",
			Name:     "Raspberry Pi B",
			Quantity: 2,
			Price:    30.00,
		}

		InventoryData := map[string]model.Item{
			item1.Sku: item1,
			item2.Sku: item2,
			item3.Sku: item3,
			item4.Sku: item4,
		}

		cartItem := model.CartOutput{
			Sku:      "A304SD",
			Name:     "",
			Quantity: 3,
			Amount:   328.5,
		}

		var CartData = model.Cart{}
		CartData.CartItems = append(CartData.CartItems, &cartItem)

		inventoryService := service.NewInventoryService(InventoryData, CartData)

		cartResult, err := inventoryService.Checkout()

		assert.Nil(t, err)
		assert.Equal(t, 295.65, cartResult.TotalPrice)
	})
}
