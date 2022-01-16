package integration__tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"todo-api/controllers"
	"todo-api/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {

	database()

	gin.SetMode(gin.TestMode)

	err := refreshTodosTable()
	if err != nil {
		log.Fatal(err)
	}
	samples := []struct {
		inputJSON  string
		statusCode int
		title      string
		body       string
		errMessage string
	}{
		{
			inputJSON:  `{"message":"buy some milk"}`,
			statusCode: 201,
			title:      "buy some milk",
			errMessage: "",
		},
		{
			inputJSON:  `{"message":"buy some milk"}`,
			statusCode: 500,
			errMessage: "title already exists",
		},
		{
			inputJSON:  `{"message":"""}`,
			statusCode: 422,
			errMessage: "Please enter a valid todo message",
		},
		{
			inputJSON:  `{"title": 12345}`,
			statusCode: 422,
			errMessage: "invalid json body",
		},
	}
	for _, v := range samples {
		r := gin.Default()
		r.POST("/todos", controllers.CreateTodo)
		req, err := http.NewRequest(http.MethodPost, "/todos", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}
		fmt.Println("this is the response data: ", responseMap)
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			//casting the interface to map:
			assert.Equal(t, responseMap["title"], v.title)
		}
		if v.statusCode == 400 || v.statusCode == 422 || v.statusCode == 500 && v.errMessage != "" {
			assert.Equal(t, responseMap["message"], v.errMessage)
		}
	}
}

func TestGetTodoByID(t *testing.T) {

	database()

	gin.SetMode(gin.TestMode)

	err := refreshTodosTable()
	if err != nil {
		log.Fatal(err)
	}
	message, err := seedOneTodo()
	if err != nil {
		t.Errorf("Error while seeding table: %s", err)
	}

	samples := []struct {
		id         string
		statusCode int
		message      string
		errMessage string
	}{
		{
			id:         strconv.Itoa(int(message.Id)),
			statusCode: 200,
			message:   message.Message,
			errMessage: "",
		},
		{
			id:         "unknwon",
			statusCode: 400,
			errMessage: "todo id should be a number",
		},
		{
			id:         strconv.Itoa(12322), //an id that does not exist
			statusCode: 404,
			errMessage: "no record matching given id",
		},
	}
	for _, v := range samples {
		r := gin.Default()
		r.GET("/todos/:todo_id", controllers.GetTodo)
		req, err := http.NewRequest(http.MethodGet, "/todos/"+v.id, nil)
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 200 {
			assert.Equal(t, responseMap["title"], v.message)
		}
		if v.statusCode == 400 || v.statusCode == 422 && v.errMessage != "" {
			assert.Equal(t, responseMap["message"], v.errMessage)
		}
	}
}

func TestGetAllTodos(t *testing.T) {

	database()

	gin.SetMode(gin.TestMode)

	err := refreshTodosTable()
	if err != nil {
		log.Fatal(err)
	}
	_, err = seedTodos()
	if err != nil {
		t.Errorf("Error while seeding table: %s", err)
	}
	r := gin.Default()
	r.GET("/todos", controllers.GetAllTodos)

	req, err := http.NewRequest(http.MethodGet, "/todos", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var todos []domain.Todo

	err = json.Unmarshal(rr.Body.Bytes(), &todos)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(todos), 2)
}