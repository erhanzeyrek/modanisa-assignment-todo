package integration__tests

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"todo-api/domain"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	queryTruncateTodo = "TRUNCATE TABLE todos;"
	queryInsertTodo  = "INSERT INTO todos(message) VALUES(?);"
	queryGetAllTodos = "SELECT id, message FROM todos;"
)
var (
	dbConn  *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("./../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	os.Exit(m.Run())
}

func database() {
	dbDriver := os.Getenv("DBDRIVER_TEST")
	username := os.Getenv("USERNAME_TEST")
	password := os.Getenv("PASSWORD_TEST")
	host := os.Getenv("HOST_TEST")
	database := os.Getenv("DATABASE_TEST")
	port := os.Getenv("PORT_TEST")

	dbConn = domain.TodoRepo.Initialize(dbDriver, username, password, port, host, database)
}

func refreshTodosTable() error {

	stmt, err := dbConn.Prepare(queryTruncateTodo)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalf("Error truncating todos table: %s", err)
	}
	return nil
}

func seedOneTodo() (domain.Todo, error) {
	todo := domain.Todo{
		Message:     "buy some milk",
	}
	stmt, err := dbConn.Prepare(queryInsertTodo)
	if err != nil {
		panic(err.Error())
	}
	insertResult, createErr := stmt.Exec(todo.Message)
	if createErr != nil {
		log.Fatalf("Error creating todo: %s", createErr)
	}
	todoId, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("Error creating todo: %s", createErr)
	}
	todo.Id = todoId
	return todo, nil
}

func seedTodos() ([]domain.Todo, error) {
	todos := []domain.Todo{
		{
			Message:     "buy some milk",
		},
		{
			Message:     "buy some chocolate",
		},
	}
	stmt, err := dbConn.Prepare(queryInsertTodo)
	if err != nil {
		panic(err.Error())
	}
	for i, _ := range todos {
		_, createErr := stmt.Exec(todos[i].Message)
		if createErr != nil {
			return nil, createErr
		}
	}
	get_stmt, err := dbConn.Prepare(queryGetAllTodos)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := get_stmt.Query()
	if err != nil {
		return nil,  err
	}
	defer rows.Close()

	results := make([]domain.Todo, 0)

	for rows.Next() {
		var todo domain.Todo
		if getError := rows.Scan(&todo.Id, &todo.Message); getError != nil {
			return nil, err
		}
		results = append(results, todo)
	}
	return results, nil
}