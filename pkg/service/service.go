package service

import (
	"github.com/redbox12/todo-app/domain"
	"github.com/redbox12/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface {
	Create(userId int, list domain.TodoList) (int, error)
	GetAll(userId int) ([]domain.TodoList, error)
	GetById(userId, listId int) (domain.TodoList, error)
	Update(userId int, listId int, input domain.UpdateListInput) error
	Delete(userId, listId int) error
}
type TodoItem interface {
	Create(userId, listId int, item domain.TodoItem) (int, error)
	GetAll(userId, listId int) ([]domain.TodoItem, error)
	GetById(userId, itemId int) (domain.TodoItem, error)
	Update(userId, itemId int, inputItem domain.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
