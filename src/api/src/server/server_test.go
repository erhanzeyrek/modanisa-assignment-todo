package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
    "models"
)

func TestPOSTTodos(t *testing.T) {
	server := &TodoServer{}

	todoMessage := `{message: "buy some milk"}`
    expectedResponse := ''

    t.Run("add todo request to store", func(t *testing.T) {
        request := addTodoRequest(todoMessage)
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)
        
        assertStatus(t, response.Code, http.StatusOK)
    })
	
}


func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Wrong status received got %d, want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Wrong response body received, got %q want %q", got, want)
	}
}

func getAllTodosRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/getAllTodos/"), nil)
	return req
}

func addTodoRequest(message string) *http.Request {
    var jsonStr = []byte(fmt.Sprintf( `{"message":"%q"}`, message))

    req, _ := http.NewRequest(http.MethodPost, "/addTodo", bytes.NewBuffer(jsonStr))
	return req
}

