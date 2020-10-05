package service

import (
	"errors"
	"fmt"

	"github.com/auliman-a/golang-test/graph/model"
	"github.com/jinzhu/copier"
)

type InventoryService interface {
	Checkout() (*model.Cart, error)
	GetInventoryData() []*model.Item
	GetCartItem() *model.Cart
	AddItemToCart(input model.CartInput) error
}

type inventoryService struct {
	inventoryData map[string]model.Item
	cartData      model.Cart
}

func NewInventoryService(invenData map[string]model.Item, cart model.Cart) InventoryService {
	return &inventoryService{
		inventoryData: invenData,
		cartData:      cart,
	}
}

func (service *inventoryService) Checkout() (*model.Cart, error) {
	err := service.ValidateQuantity(service.cartData)
	if err != nil {
		return nil, err
	}
	service.CheckAndApplyPromotion()
	service.CalculateTotalAmount()
	service.ProcessInventory()
	modelCartDataCopy := service.cartData
	service.cartData = model.Cart{}
	return &modelCartDataCopy, nil
}

// func (r *queryResolver) GetCartItems(ctx context.Context) (*model.Cart, error) {
// 	CheckAndApplyPromotion()
// 	return &model.CartData, nil
// }

func (service *inventoryService) ValidateQuantity(cart model.Cart) error {
	for i := range cart.CartItems {
		invenData := service.inventoryData[cart.CartItems[i].Sku]
		if cart.CartItems[i].Quantity > invenData.Quantity {
			return errors.New("Not Enough Quantity for Item")
		}
	}
	return nil
}

func (service *inventoryService) ProcessInventory() {
	for i := range service.cartData.CartItems {
		invenData, _ := service.inventoryData[service.cartData.CartItems[i].Sku]
		invenData.Quantity -= service.cartData.CartItems[i].Quantity
		service.inventoryData[service.cartData.CartItems[i].Sku] = invenData
	}
}

func (service *inventoryService) GetInventoryData() []*model.Item {
	var inventoryData []*model.Item

	for _, value := range service.inventoryData {
		var item model.Item
		copier.Copy(&item, &value)
		inventoryData = append(inventoryData, &item)
	}
	return inventoryData
}

func (service *inventoryService) GetCartItem() *model.Cart {
	return &service.cartData
}

func (service *inventoryService) AddItemToCart(input model.CartInput) error {
	invenData := model.InventoryData[input.Sku]

	if invenData != (model.Item{}) {

		existCartItem := service.getCartItem(input.Sku)
		if existCartItem != nil {
			existCartItem.Quantity += input.Quantity
			existCartItem.Amount += float64(input.Quantity) * invenData.Price
		} else {
			cartItem := model.CartOutput{
				Sku:      invenData.Sku,
				Name:     invenData.Name,
				Quantity: input.Quantity,
				Amount:   float64(input.Quantity) * invenData.Price,
			}
			service.cartData.CartItems = append(service.cartData.CartItems, &cartItem)
		}
	} else {
		return errors.New("Can not find item")
	}
	return nil
}
func (service *inventoryService) CheckAndApplyPromotion() {
	var cartCopies []model.CartOutput

	for i := range service.cartData.CartItems {
		var cartCopy model.CartOutput

		copier.Copy(&cartCopy, &service.cartData.CartItems[i])
		cartCopies = append(cartCopies, cartCopy)
	}

	for i := range cartCopies {
		if cartCopies[i].Sku == "43N23P" {
			existCartItem := service.getCartItem("234234")
			if existCartItem == nil {
				cartItem := model.CartOutput{
					Sku:      "234234",
					Name:     model.InventoryData["234234"].Name,
					Quantity: cartCopies[i].Quantity,
					Amount:   0,
				}
				service.cartData.CartItems = append(service.cartData.CartItems, &cartItem)
			} else {
				if existCartItem.Quantity <= cartCopies[i].Quantity {
					existCartItem.Amount = 0
				} else {
					existCartItem.Amount = model.InventoryData["234234"].Price * (float64(existCartItem.Quantity - cartCopies[i].Quantity))
				}
			}
		}

		if cartCopies[i].Sku == "120P90" {
			existCartItem := service.getCartItem("120P90")
			existCartItem.Amount = model.InventoryData["120P90"].Price * float64(existCartItem.Quantity-(existCartItem.Quantity/3))
		}

		if cartCopies[i].Sku == "A304SD" {
			if cartCopies[i].Quantity >= 3 {
				existCartItem := service.getCartItem("A304SD")
				amount := existCartItem.Amount - (existCartItem.Amount * 10 / 100)
				existCartItem.Amount = amount
				fmt.Println(existCartItem)
			}
		}
	}
}
func (service *inventoryService) CalculateTotalAmount() {
	for i := range service.cartData.CartItems {
		service.cartData.TotalPrice += service.cartData.CartItems[i].Amount
	}
}

func (service *inventoryService) getCartItem(sku string) *model.CartOutput {
	for i := range service.cartData.CartItems {
		if service.cartData.CartItems[i].Sku == sku {
			// Found!
			return service.cartData.CartItems[i]
		}
	}
	return nil
}
