package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/auliman-a/golang-test/graph/generated"
	"github.com/auliman-a/golang-test/graph/model"
	"github.com/auliman-a/golang-test/service"
)

var (
	_inventoryService service.InventoryService = service.NewInventoryService(model.InventoryData, model.CartData)
)

func (r *mutationResolver) AddToCart(ctx context.Context, input model.CartInput) (string, error) {
	err := _inventoryService.AddItemToCart(input)

	if err != nil {
		return "Error", err
	}
	return "Successfuly adding item to cart", nil
}

func (r *mutationResolver) Checkout(ctx context.Context) (*model.Cart, error) {
	cartCheck, err := _inventoryService.Checkout()
	if err != nil {
		return nil, err
	}
	return cartCheck, nil
}

func (r *queryResolver) GetInventoryData(ctx context.Context) ([]*model.Item, error) {
	inventoryData := _inventoryService.GetInventoryData()
	return inventoryData, nil
}

func (r *queryResolver) GetCartItem(ctx context.Context) (*model.Cart, error) {
	return _inventoryService.GetCartItem(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
