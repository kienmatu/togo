package todos

import (
	"context"
	"dangquang9a/go-location/internal/models"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *models.Todo) error
	GetTodosByUserId(ctx context.Context, userId string) ([]*models.Todo, error)
	GetAllTodos(ctx context.Context) ([]*models.Todo, error)
	CountTodo(ctx context.Context, userId string) (int, error)
}
