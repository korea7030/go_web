package model

import (
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type dbHandler interface {
	GetTodos() []*Todo
	addTodo(name string) *Todo
	RemoveTodo(id int) bool
	CompleteTodo(id int, complete bool) bool
}

var handler dbHandler

func init() {
	// handler = newMemoryHandler()
	handler = newSqliteHandler()
}

func newMemoryHandler() dbHandler {
	m := &memoryHandler{}
	m.todoMap = make(map[int]*Todo)
	return m
}

func GetTodos() []*Todo {
	return handler.GetTodos()
}

func AddTodo(name string) *Todo {
	return handler.addTodo(name)
}

func RemoveTodo(id int) bool {
	return handler.RemoveTodo(id)
}

func CompleteTodo(id int, complete bool) bool {
	return handler.CompleteTodo(id, complete)
}
