package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

type TodoStore interface {
	GetAllTodos() []models.Todo
	AddTodo(models.Todo)
}

type TodoServer struct {
	store TodoStore
}

func (p *TodoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	endpoint := path.Base(r.URL.Path)

	switch endpoint {
	case "addTodo":
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
	
		if err != nil {
			log.Fatalln(err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
	
		var todo models.Todo
		json.Unmarshal(body, &todo)
		p.addTodo(w, todo)
	case "getAllTodos":
		p.getAllTodos(w)
	}
}

func (p *TodoServer) getAllTodos(w http.ResponseWriter) {
	todos := p.store.GetAllTodos()

	fmt.Fprint(w, todos)
}

func (p *TodoServer) addTodo(w http.ResponseWriter, todo models.Todo) {
	p.store.AddTodo(todo)
	w.WriteHeader(http.StatusOK)
}