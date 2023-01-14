package usecase

import (
	"context"
	"dangquang9a/go-location/internal/auth"
	"dangquang9a/go-location/internal/models"
	"dangquang9a/go-location/internal/todos"
	"time"

	"github.com/google/uuid"
)

type todoUsecase struct {
	todoRepo todos.TodoRepository
	userRepo auth.UserRepository
}

func NewTodoUseCase(todoRepo todos.TodoRepository, userRepo auth.UserRepository) todos.UseCase {
	return &todoUsecase{
		todoRepo: todoRepo,
		userRepo: userRepo,
	}
}

func (tu todoUsecase) CreateTodo(ctx context.Context, userId string, content string) error {
	todo := &models.Location{
		Id:        uuid.New().String(),
		Name:      content,
		CreatedAt: time.Now(),
		CreatedBy: userId,
	}

	return tu.todoRepo.CreateTodo(ctx, todo)

}

func (tu todoUsecase) GetTodosByUserId(ctx context.Context, userId string) ([]*models.Location, error) {
	todos, err := tu.todoRepo.GetTodosByUserId(ctx, userId)

	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (tu todoUsecase) GetAllTodos(ctx context.Context) ([]*models.Location, error) {
	todos, err := tu.todoRepo.GetAllTodos(ctx)

	if err != nil {
		return nil, err
	}
	return todos, nil
}
