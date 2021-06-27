package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *mutationResolver) User(ctx context.Context) (*model.UserOps, error) {
	return &model.UserOps{}, nil
}

func (r *mutationResolver) Item(ctx context.Context) (*model.ItemOps, error) {
	return &model.ItemOps{}, nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	return service.UserGetByID(ctx, id)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return service.UserGetAll(ctx)
}

func (r *queryResolver) Item(ctx context.Context, id int) (*model.Item, error) {
	return service.ItemGetByID(ctx, id)
}

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	return service.ItemGetAll(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
