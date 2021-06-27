package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/dataloader"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *userResolver) Items(ctx context.Context, obj *model.User) ([]*model.Item, error) {
	return dataloader.For(ctx).ItemByUserID.Load(obj.ID)
}

func (r *userOpsResolver) Create(ctx context.Context, obj *model.UserOps, input model.NewUser) (*model.User, error) {
	return service.UserCreate(ctx, input)
}

func (r *userOpsResolver) Update(ctx context.Context, obj *model.UserOps, input model.UpdateUser) (*string, error) {
	return service.UserUpdate(ctx, input)
}

func (r *userOpsResolver) Delete(ctx context.Context, obj *model.UserOps, id int) (*string, error) {
	return service.UserDelete(ctx, id)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// UserOps returns generated.UserOpsResolver implementation.
func (r *Resolver) UserOps() generated.UserOpsResolver { return &userOpsResolver{r} }

type userResolver struct{ *Resolver }
type userOpsResolver struct{ *Resolver }
