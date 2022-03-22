package todos

import (
	"context"
	"kienmatu/go-todos/internal/models"
)

type UseCase interface {
	CreateTodo(ctx context.Context, userId string, content string) error
	GetTodosByUserId(ctx context.Context, userId string) ([]*models.Todo, error)
	GetAllTodos(ctx context.Context) ([]*models.Todo, error)
}
