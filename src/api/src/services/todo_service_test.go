package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"testing"
	"time"
	"todo-api/domain"
	"todo-api/utils/error_utils"

	"github.com/stretchr/testify/assert"
)

var (
	tm = time.Now()
	getTodoDomain func(todoId int64) (*domain.Todo, error_utils.TodoErr)
	createTodoDomain func(todo *domain.Todo) (*domain.Todo, error_utils.TodoErr)
	getAllTodosDomain func() ([]domain.Todo, error_utils.TodoErr)
)

type getDBMock struct {}

func (m *getDBMock) Get(todoId int64) (*domain.Todo, error_utils.TodoErr){
	return getTodoDomain(todoId)
}
func (m *getDBMock) Create(todo *domain.Todo) (*domain.Todo, error_utils.TodoErr){
	return createTodoDomain(todo)
}
func (m *getDBMock) GetAll() ([]domain.Todo, error_utils.TodoErr) {
	return getAllTodosDomain()
}
func (m *getDBMock) Initialize(string, string, string, string, string, string) *sql.DB  {
	return nil
}


///////////////////////////////////////////////////////////////
// Start of "GetTodo" test cases
///////////////////////////////////////////////////////////////
func TestMessagesService_GetMessage_Success(t *testing.T) {
	domain.TodoRepo = &getDBMock{} //this is where we swapped the functionality
	getTodoDomain = func(todoId int64) (*domain.Todo, error_utils.TodoErr) {
		return &domain.Todo{
			Id:        1,
			Message:     "buy some milk",
		}, nil
	}
	todo, err := TodoService.GetTodo(1)
	fmt.Println("this is the message: ", todo)
	assert.NotNil(t, todo)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, todo.Id)
	assert.EqualValues(t, "buy some milk", todo.Message)
}

func TestMessagesService_GetMessage_NotFoundID(t *testing.T) {
	domain.TodoRepo = &getDBMock{}
	//TodoService = &serviceMock{}

	getTodoDomain = func(todoId int64) (*domain.Todo, error_utils.TodoErr) {
		return nil, error_utils.NewNotFoundError("the todo id is not found")
	}
	todo, err := TodoService.GetTodo(1)
	assert.Nil(t, todo)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "the todo id is not found", err.Message())
	assert.EqualValues(t, "not_found", err.Error())
}
///////////////////////////////////////////////////////////////
// End of "GetTodo" test cases
///////////////////////////////////////////////////////////////


///////////////////////////////////////////////////////////////
// Start of	"CreateTodo" test cases
///////////////////////////////////////////////////////////////

func TestMessagesService_CreateTodo_Success(t *testing.T) {
	domain.TodoRepo = &getDBMock{}
	createTodoDomain  = func(todo *domain.Todo) (*domain.Todo, error_utils.TodoErr){
		return &domain.Todo{
			Id:        1,
			Message:   "buy some milk",
		}, nil
	}
	request := &domain.Todo{
		Message:     "buy some milk",
	}
	todo, err := TodoService.CreateTodo(request)
	fmt.Println("this is the todo: ", todo)
	assert.NotNil(t, todo)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, todo.Id)
	assert.EqualValues(t, "buy some milk", todo.Message)
}

func TestMessagesService_CreateTodo_Invalid_Request(t *testing.T) {
	tests := []struct {
		request *domain.Todo
		statusCode int
		errMsg string
		errErr string
	}{
		{
			request: &domain.Todo{
			  Message:     "",
		    },
		    statusCode: http.StatusUnprocessableEntity,
		    errMsg: "Please enter a valid todo message",
		    errErr: "invalid_request",
		},		
	}
	for _, tt := range tests {
		todo, err := TodoService.CreateTodo(tt.request)
		assert.Nil(t, todo)
		assert.NotNil(t, err)
		assert.EqualValues(t, tt.errMsg, err.Message())
		assert.EqualValues(t, tt.statusCode, err.Status())
		assert.EqualValues(t, tt.errErr, err.Error())
	}
}

func TestMessagesService_CreateTodo_Failure(t *testing.T) {
	domain.TodoRepo = &getDBMock{}
	createTodoDomain  = func(todo *domain.Todo) (*domain.Todo, error_utils.TodoErr){
		return nil, error_utils.NewInternalServerError("todo is already exists")
	}
	request := &domain.Todo{
		Message:     "buy some milk",
	}
	todo, err := TodoService.CreateTodo(request)
	assert.Nil(t, todo)
	assert.NotNil(t, err)
	assert.EqualValues(t, "todo is already exists", err.Message())
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "server_error", err.Error())
}

///////////////////////////////////////////////////////////////
// End of "CreateTodo" test cases
///////////////////////////////////////////////////////////////


///////////////////////////////////////////////////////////////
// Start of "GetAllTodos" test cases
///////////////////////////////////////////////////////////////
func TestMessagesService_GetAllMessages(t *testing.T) {
	domain.TodoRepo = &getDBMock{}
	getAllTodosDomain  = func() ([]domain.Todo, error_utils.TodoErr) {
		return []domain.Todo{
			{
				Id:        1,
				Message:     "buy some milk",
			},
			{
				Id:        2,
				Message:     "buy some chocolate",
			},
		}, nil
	}
	messages, err := TodoService.GetAllTodos()
	assert.Nil(t, err)
	assert.NotNil(t, messages)
	assert.EqualValues(t, messages[0].Id, 1)
	assert.EqualValues(t, messages[0].Message, "buy some milk")
	assert.EqualValues(t, messages[1].Id, 2)
	assert.EqualValues(t, messages[1].Message, "buy some chocolate")
}

func TestMessagesService_GetAllMessages_Error_Getting_Messages(t *testing.T) {
	domain.TodoRepo = &getDBMock{}
	getAllTodosDomain  = func() ([]domain.Todo, error_utils.TodoErr) {
		return nil, error_utils.NewInternalServerError("error getting todos")
	}
	messages, err := TodoService.GetAllTodos()
	assert.NotNil(t, err)
	assert.Nil(t, messages)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "error getting todos", err.Message())
	assert.EqualValues(t, "server_error", err.Error())
}
///////////////////////////////////////////////////////////////
// End of "GetAllMessage" test cases
///////////////////////////////////////////////////////////////