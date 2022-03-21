package repository

import (
	"context"
	"kienmatu/go-todos/internal/models"
	"kienmatu/go-todos/internal/todos"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) todos.TodoRepository {
	return &todoRepository{db: db}
}

func (tr *todoRepository) CreateTodo(ctx context.Context, todo *models.Todo) error {
	result := tr.db.Create(&todo)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (tr *todoRepository) GetTodosByUserId(ctx context.Context, userId string) ([]*models.Todo, error) {
	var todos []*models.Todo
	err := tr.db.Where(&models.Todo{CreatedBy: userId}).Find(&todos).Error

	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (tr *todoRepository) GetAllTodos(ctx context.Context) ([]*models.Todo, error) {
	var todos []*models.Todo
	// can add offset later
	err := tr.db.Limit(200).Find(&todos).Error

	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (tr *todoRepository) CountTodo(ctx context.Context, userId string) (int, error) {
	// tr.db.Raw(`SELECT
	// 		DATE_TRUNC('day', "createdAt") AS "alias1",
	// 		COUNT("createdAt") AS "alias2"
	// 		FROM "todos"
	// 		GROUP BY DATE_TRUNC('day', "createdAt");`)
	return 1, nil
}
