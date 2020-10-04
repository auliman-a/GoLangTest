package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/auliman-a/golang-test/graph/generated"
	"github.com/auliman-a/golang-test/graph/model"
	"github.com/jinzhu/copier"
)

func (r *mutationResolver) AddToCart(ctx context.Context, input model.CartInput) (string, error) {
	err := AddItemToCart(input)

	if err != nil {
		return "", err
	}
	return "Successfuly adding item to cart", nil
}

func (r *mutationResolver) Checkout(ctx context.Context, input []*model.CheckoutInput) ([]*model.Cart, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCartItems(ctx context.Context) (*model.Cart, error) {
	CheckAndApplyPromotion()
	return &model.CartData, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func AddItemToCart(input model.CartInput) error {
	invenData := model.InventoryData[input.Sku]

	if invenData != (model.Item{}) {

		existCartItem := GetCartItem(input.Sku)
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
			model.CartData.CartItems = append(model.CartData.CartItems, &cartItem)
		}
	} else {
		return errors.New("Can not find item")
	}
	return nil
}

func CheckAndApplyPromotion() {
	var cartCopies []model.CartOutput

	for i := range model.CartData.CartItems {
		var cartCopy model.CartOutput

		copier.Copy(&cartCopy, &model.CartData.CartItems[i])
		cartCopies = append(cartCopies, cartCopy)
	}

	for i := range cartCopies {
		if cartCopies[i].Sku == "43N23P" {
			existCartItem := GetCartItem("234234")
			if existCartItem == nil {
				cartItem := model.CartOutput{
					Sku:      "234234",
					Name:     model.InventoryData["234234"].Name,
					Quantity: cartCopies[i].Quantity,
					Amount:   0,
				}
				model.CartData.CartItems = append(model.CartData.CartItems, &cartItem)
			} else {
				if existCartItem.Quantity <= cartCopies[i].Quantity {
					existCartItem.Amount = 0
				} else {
					existCartItem.Amount = model.InventoryData["234234"].Price * (float64(existCartItem.Quantity - cartCopies[i].Quantity))
				}
			}
		}

		if cartCopies[i].Sku == "120P90" {
			existCartItem := GetCartItem("120P90")
			existCartItem.Amount = model.InventoryData["120P90"].Price * float64(existCartItem.Quantity-(existCartItem.Quantity/3))
		}

		if cartCopies[i].Sku == "A304SD" {
			if cartCopies[i].Quantity >= 3 {
				existCartItem := GetCartItem("A304SD")
				amount := existCartItem.Amount - (existCartItem.Amount * 10 / 100)
				existCartItem.Amount = amount
				fmt.Println(existCartItem)
			}
		}
	}

	//Applying promotion
	// if sku == "43N23P" {
	// 	invenData := model.InventoryData["234234"]
	// 	cartItem := model.CartOutput{
	// 		Sku:      invenData.Sku,
	// 		Name:     invenData.Name,
	// 		Quantity: quantity,
	// 		Amount:   0,
	// 	}
	// 	model.CartData.CartItems = append(model.CartData.CartItems, &cartItem)
	// }
	// if sku == "43N23P" {

	// }
}

func GetCartItem(sku string) *model.CartOutput {
	for i := range model.CartData.CartItems {
		if model.CartData.CartItems[i].Sku == sku {
			// Found!
			return model.CartData.CartItems[i]
		}
	}
	return nil
}
