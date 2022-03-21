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
	result := tr.db.WithContext(ctx).Create(&todo)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tr *todoRepository) GetTodosByUserId(ctx context.Context, userId string) ([]*models.Todo, error) {
	var todos []*models.Todo
	err := tr.db.WithContext(ctx).Where(&models.Todo{CreatedBy: userId}).Find(&todos).Error

	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (tr *todoRepository) GetAllTodos(ctx context.Context) ([]*models.Todo, error) {
	var todos []*models.Todo
	// can add offset later
	err := tr.db.WithContext(ctx).Limit(200).Find(&todos).Error

	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (tr *todoRepository) CountTodo(ctx context.Context, userId string) (int, error) {
	var count int
	err := tr.db.WithContext(ctx).Raw(`SELECT 
			COUNT(*)
			FROM "todos"
			WHERE todos.created_by = ?
			AND DATE_TRUNC('day', "created_at") = CURRENT_DATE
			GROUP BY DATE_TRUNC('day', "created_at")`, userId).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
