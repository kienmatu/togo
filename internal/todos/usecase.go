package todos

import (
	"context"
	"dangquang9a/go-location/internal/models"
)

type UseCase interface {
	CreateTodo(ctx context.Context, userId string, content string) error
	GetTodosByUserId(ctx context.Context, userId string) ([]*models.Location, error)
	GetAllTodos(ctx context.Context) ([]*models.Location, error)
}
