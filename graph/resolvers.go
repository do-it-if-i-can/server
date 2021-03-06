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

// query  ----------------------------------------------

func (r *queryResolver) GetUserByID(ctx context.Context, input model.GetUserByID) (*model.User, error) {
	user := &model.User{}
	todos := []model.Todo{}
	if err := r.DB.First(&user, "id = ?", input.UserID).Error; err != nil {
		return nil, err
	}
	if err := r.DB.Find(&todos, "user_id = ?", input.UserID).Error; err != nil {
		return nil, err
	}
	user.Todos = todos
	return user, nil
}

func (r *queryResolver) GetTodosByCategory(ctx context.Context, input model.GetTodosByCategory) ([]*model.Todo, error) {
	todos := []*model.Todo{}
	if err := r.DB.Where(&model.Todo{UserID: input.UserID, Category: input.Category}).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *queryResolver) GetTodosByUser(ctx context.Context, input model.GetTodosByUser) ([]*model.Todo, error) {
	return nil, nil
}

// mutation ----------------------------------------------

func (r *mutationResolver) UpsertUser(ctx context.Context, input model.UpsertUser) (*model.User, error) {
	return nil, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (bool, error) {
	todo := &model.Todo{
		Title:    input.Title,
		Category: input.Category,
		UserID:   input.UserID,
	}
	if err := r.DB.Create(&todo).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (bool, error) {
	return true, nil
}

func (r *mutationResolver) UpdateTodoDone(ctx context.Context, input model.UpdateTodoDone) (bool, error) {
	return true, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.DeleteTodo) (bool, error) {
	return true, nil
}

func (r *mutationResolver) CopyTodo(ctx context.Context, input model.CopyTodo) (bool, error) {
	return true, nil
}

func (r *mutationResolver) MoveTodo(ctx context.Context, input model.MoveTodo) (bool, error) {
	return true, nil
}
