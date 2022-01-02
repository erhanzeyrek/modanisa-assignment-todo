package domain

import (
	"database/sql"
	"fmt"
	"log"
	"todo-api/utils/error_formats"
	"todo-api/utils/error_utils"

	_ "github.com/go-sql-driver/mysql"
)

var (
	TodoRepo todoRepoInterface = &todoeRepo{}
)

const (
	queryGetMessage    = "SELECT id, message FROM todos WHERE id=?;"
	queryInsertMessage = "INSERT INTO todos(message) VALUES(?);"
	queryGetAllMessages = "SELECT id, message FROM todos;"
)

type todoRepoInterface interface {
	Get(int64) (*Todo, error_utils.TodoErr)
	Create(*Todo) (*Todo, error_utils.TodoErr)
	GetAll() ([]Todo, error_utils.TodoErr)
	Initialize(string, string, string, string, string, string) *sql.DB
}
type todoeRepo struct {
	db *sql.DB
}

func (mr *todoeRepo) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *sql.DB  {
	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	mr.db, err = sql.Open(Dbdriver, DBURL)
	if err != nil {
		log.Fatal("This is the error connecting to the database:", err)
	}
	fmt.Printf("We are connected to the %s database", Dbdriver)

	return mr.db
}

func NewTodoRepository(db *sql.DB) todoRepoInterface {
	return &todoeRepo{db: db}
}

func (mr *todoeRepo) Get(messageId int64) (*Todo, error_utils.TodoErr) {
	stmt, err := mr.db.Prepare(queryGetMessage)
	if err != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("Error when trying to prepare todo: %s", err.Error()))
	}
	defer stmt.Close()

	var todo Todo
	result := stmt.QueryRow(messageId)
	if getError := result.Scan(&todo.Id, &todo.Message); getError != nil {
		fmt.Println("Error when trying to get todo: ", getError)
		return nil,  error_formats.ParseError(getError)
	}
	return &todo, nil
}

func (mr *todoeRepo) GetAll() ([]Todo, error_utils.TodoErr) {
	stmt, err := mr.db.Prepare(queryGetAllMessages)
	if err != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("Error when trying to prepare all todos: %s", err.Error()))
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil,  error_formats.ParseError(err)
	}
	defer rows.Close()

	results := make([]Todo, 0)

	for rows.Next() {
		var todo Todo
		if getError := rows.Scan(&todo.Id, &todo.Message); getError != nil {
			return nil, error_utils.NewInternalServerError(fmt.Sprintf("Error when trying to get todo: %s", getError.Error()))
		}
		results = append(results, todo)
	}
	if len(results) == 0 {
		return nil, error_utils.NewNotFoundError("no records found")
	}
	return results, nil
}

func (mr *todoeRepo) Create(msg *Todo) (*Todo, error_utils.TodoErr) {
	fmt.Println("WE REACHED THE DOMAIN")
	stmt, err := mr.db.Prepare(queryInsertMessage)
	if err != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("error when trying to prepare todo to save: %s", err.Error()))
	}
	fmt.Println("WE DIDNT REACH HERE")

	defer stmt.Close()

	insertResult, createErr := stmt.Exec(msg.Message)
	if createErr != nil {
		return nil,  error_formats.ParseError(createErr)
	}
	msgId, err := insertResult.LastInsertId()
	if err != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("error when trying to save Todo: %s", err.Error()))
	}
	msg.Id = msgId

	return msg, nil
}