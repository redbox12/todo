package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/redbox12/todo-app/internal/domain"
)

const (
	usersTable      = "users"
	todoListTable   = "todo_lists"
	usersListTable  = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)


type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}
type TodoList interface {
	Create(userId int, list domain.TodoList) (int, error)
	GetAll(userId int) ([]domain.TodoList, error)
	GetById(userId, listId int) (domain.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input domain.UpdateListInput) error
}
type TodoItem interface {
	Create(listId int, item domain.TodoItem) (int, error)
	GetAll(userId, listId int) ([]domain.TodoItem, error)
	GetById(userId, itemId int) (domain.TodoItem, error)
	Update(userId, itemId int, inputItem domain.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
