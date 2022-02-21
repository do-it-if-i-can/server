package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/do-it-if-i-can/server/graph/generated"
	"github.com/do-it-if-i-can/server/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Title:  input.Title,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID, // fix this line
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context, category *model.Category) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *todoResolver) UpdatedAt(ctx context.Context, obj *model.Todo) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) CreatedAt(ctx context.Context, obj *model.Todo) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) UpdatedAt(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) CreatedAt(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *todoResolver) Category(ctx context.Context, obj *model.Todo) (model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *todoResolver) Priority(ctx context.Context, obj *model.Todo) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *todoResolver) Title(ctx context.Context, obj *model.Todo) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *todoResolver) Description(ctx context.Context, obj *model.Todo) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) Avatar(ctx context.Context, obj *model.User) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}
