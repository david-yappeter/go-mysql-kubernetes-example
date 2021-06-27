package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *itemOpsResolver) Create(ctx context.Context, obj *model.ItemOps, input model.NewItem) (*model.Item, error) {
	return service.ItemCreate(ctx, input)
}

func (r *itemOpsResolver) Update(ctx context.Context, obj *model.ItemOps, input model.UpdateItem) (*string, error) {
	return service.ItemUpdate(ctx, input)
}

func (r *itemOpsResolver) Delete(ctx context.Context, obj *model.ItemOps, id int) (*string, error) {
	return service.ItemDelete(ctx, id)
}

// ItemOps returns generated.ItemOpsResolver implementation.
func (r *Resolver) ItemOps() generated.ItemOpsResolver { return &itemOpsResolver{r} }

type itemOpsResolver struct{ *Resolver }
