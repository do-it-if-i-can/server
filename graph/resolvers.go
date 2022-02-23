package graph

import (
	"context"

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
		Title:    input.Title,
		Category: input.Category,
		UserID:   input.UserID,
	}
	if err := r.DB.Create(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

// query  ----------------------------------------------

func (r *queryResolver) GetTodosByCategory(ctx context.Context, input model.GetTodosByCategory) ([]*model.Todo, error) {
	todos := []*model.Todo{}
	if err := r.DB.Where(&model.Todo{UserID: input.UserID, Category: input.Category}).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
