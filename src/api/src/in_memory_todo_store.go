package main

import (
	"sync"
	todo "todo-server/models"
)
func NewInMemoryTodoStore() *InMemoryTodoStore {
	return &InMemoryTodoStore{
		[]todo.Todo{},
		sync.RWMutex{},
	}
}

type InMemoryTodoStore struct {
	store []todo.Todo
	lock sync.RWMutex
}

func (i *InMemoryTodoStore) AddTodo(todo todo.Todo) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.store[todo]++
}


func (i *InMemoryTodoStore) GetAllTodos() int {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.store
}
