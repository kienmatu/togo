package usecase

import (
	"context"
	"errors"
	"kienmatu/go-todos/internal/auth"
	"kienmatu/go-todos/internal/models"
	"kienmatu/go-todos/internal/todos"
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
	todo := &models.Todo{
		Id:        uuid.New().String(),
		Content:   content,
		CreatedAt: time.Now(),
		CreatedBy: userId,
	}
	count, err := tu.todoRepo.CountTodo(ctx, userId)
	if err != nil {
		return err
	}
	user, err := tu.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return err
	}
	if user.Limit > count {
		return tu.todoRepo.CreateTodo(ctx, todo)
	} else {
		return errors.New("limit exceeded")
	}
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
