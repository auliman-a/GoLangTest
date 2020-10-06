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
	t.Run("Test CheckOut with Macbook Promo", func(t *testing.T) {

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

	t.Run("Test CheckOut with Google Home Promo ", func(t *testing.T) {

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

	t.Run("Test CheckOut with Alexa Speaker Promo ", func(t *testing.T) {

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

	t.Run("Test CheckOut failed Quantity not enough", func(t *testing.T) {

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
			Sku:      "234234",
			Name:     "",
			Quantity: 3,
			Amount:   90.00,
		}

		var CartData = model.Cart{}
		CartData.CartItems = append(CartData.CartItems, &cartItem)

		inventoryService := service.NewInventoryService(InventoryData, CartData)

		cartResult, err := inventoryService.Checkout()

		assert.Nil(t, cartResult)
		assert.Equal(t, "Not Enough Quantity for Item", err.Error())
	})

	t.Run("Test Add Item to Cart", func(t *testing.T) {

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

		cartInput := model.CartInput{
			Sku:      "234234",
			Quantity: 3,
		}

		var CartData = model.Cart{}

		inventoryService := service.NewInventoryService(InventoryData, CartData)

		err := inventoryService.AddItemToCart(cartInput)

		assert.Nil(t, err)
	})

	t.Run("Test Add Item to Cart with existing Item in Cart", func(t *testing.T) {

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

		cartInput := model.CartInput{
			Sku:      "120P90",
			Quantity: 3,
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

		err := inventoryService.AddItemToCart(cartInput)

		assert.Nil(t, err)
		assert.Equal(t, 6, CartData.CartItems[0].Quantity)
	})

	t.Run("Test Add Item to Cart - Missing Inventory Data", func(t *testing.T) {

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

		cartInput := model.CartInput{
			Sku:      "X12345",
			Quantity: 3,
		}

		var CartData = model.Cart{}

		inventoryService := service.NewInventoryService(InventoryData, CartData)

		err := inventoryService.AddItemToCart(cartInput)

		assert.NotNil(t, err)
		assert.Equal(t, "Can not find item", err.Error())
	})

	t.Run("Test Get Inventory Data", func(t *testing.T) {

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

		var CartData = model.Cart{}

		inventoryService := service.NewInventoryService(InventoryData, CartData)

		inventoryData := inventoryService.GetInventoryData()

		assert.NotNil(t, inventoryData)
		assert.Equal(t, 4, len(inventoryData))
	})
}
