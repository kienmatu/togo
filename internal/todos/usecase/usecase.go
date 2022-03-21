package usecase

import (
	"context"
	"kienmatu/go-todos/internal/models"
	"kienmatu/go-todos/internal/todos"
	"time"

	"github.com/google/uuid"
)

type todoUsecase struct {
	todoRepo todos.TodoRepository
}

func NewTodoUseCase(todoRepo todos.TodoRepository) todos.UseCase {
	return &todoUsecase{
		todoRepo: todoRepo,
	}
}

func (tu todoUsecase) CreateTodo(ctx context.Context, userId string, content string) error {
	todo := &models.Todo{
		Id:        uuid.New().String(),
		Content:   content,
		CreatedAt: time.Now(),
		CreatedBy: userId,
	}
	return tu.todoRepo.CreateTodo(ctx, todo)
}

func (tu todoUsecase) GetTodosByUserId(ctx context.Context, userId string) ([]*models.Todo, error) {
	todos, err := tu.todoRepo.GetTodosByUserId(ctx, userId)

	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (tu todoUsecase) GetAllTodos(ctx context.Context) ([]*models.Todo, error) {
	todos, err := tu.todoRepo.GetAllTodos(ctx)

	if err != nil {
		return nil, err
	}
	return todos, nil
}
