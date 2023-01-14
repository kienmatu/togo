package todos

import (
	"context"
	"dangquang9a/go-location/internal/models"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *models.Location) error
	GetTodosByUserId(ctx context.Context, userId string) ([]*models.Location, error)
	GetAllTodos(ctx context.Context) ([]*models.Location, error)
	CountTodo(ctx context.Context, userId string) (int, error)
}
