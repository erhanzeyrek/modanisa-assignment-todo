package services

import (
	"todo-api/domain"
	"todo-api/utils/error_utils"
)

var (
	TodoService todoServiceInterface = &todoService{}
)

type todoService struct{}

type todoServiceInterface interface {
	GetTodo(int64) (*domain.Todo, error_utils.TodoErr)
	CreateTodo(*domain.Todo) (*domain.Todo, error_utils.TodoErr)
	GetAllTodos() ([]domain.Todo, error_utils.TodoErr)
}

func (m *todoService) GetTodo(msgId int64) (*domain.Todo, error_utils.TodoErr) {
	todo, err := domain.TodoRepo.Get(msgId)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (m *todoService) GetAllTodos() ([]domain.Todo, error_utils.TodoErr) {
	messages, err := domain.TodoRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (m *todoService) CreateTodo(todo *domain.Todo) (*domain.Todo, error_utils.TodoErr) {
	if err := todo.Validate(); err != nil {
		return nil, err
	}
	todo, err := domain.TodoRepo.Create(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}