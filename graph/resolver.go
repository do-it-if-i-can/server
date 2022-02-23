package graph

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/do-it-if-i-can/server/graph/generated"
	"github.com/do-it-if-i-can/server/graph/model"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *gorm.DB
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// mutation ----------------------------------------------

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Title:  input.Title,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	fmt.Println("completed dummy create todo.")
	return todo, nil
}

// query  ----------------------------------------------

func (r *queryResolver) Todos(ctx context.Context, category *model.Category) ([]*model.Todo, error) {
	array := []*model.Todo{}
	return array, nil
}
