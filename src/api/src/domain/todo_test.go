package domain

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

var created_at = time.Now()

func TestTodoRepo_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := NewTodoRepository(db)

	tests := []struct {
		name    string
		s       todoRepoInterface
		msgId   int64
		mock    func()
		want    *Todo
		wantErr bool
	}{
		{
			//When everything works as expected
			name:  "OK",
			s:     s,
			msgId: 1,
			mock: func() {
				//We added one row
				rows := sqlmock.NewRows([]string{"Id", "Message"}).AddRow(1, "buy some milk")
				mock.ExpectPrepare("SELECT (.+) FROM todos").ExpectQuery().WithArgs(1).WillReturnRows(rows)
			},
			want: &Todo{
				Id:        1,
				Message: "buy some milk",
			},
		},
		{
			name:  "Not Found",
			s:     s,
			msgId: 1,
			mock: func() {
				rows := sqlmock.NewRows([]string{"Id", "Message"})
				mock.ExpectPrepare("SELECT (.+) FROM todos").ExpectQuery().WithArgs(1).WillReturnRows(rows)
			},
			wantErr: true,
		},
		{
			name:  "Invalid Prepare",
			s:     s,
			msgId: 1,
			mock: func() {
				rows := sqlmock.NewRows([]string{"Id", "Message"}).AddRow(1, "buy some milk")
				mock.ExpectPrepare("SELECT (.+) FROM wrong_table").ExpectQuery().WithArgs(1).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Get(tt.msgId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v, message %v", err, tt.wantErr, err.Message())
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database", err)
	}
	defer db.Close()
	s := NewTodoRepository(db)

	tests := []struct {
		name    string
		s       todoRepoInterface
		request *Todo
		mock    func()
		want    *Todo
		wantErr bool
	}{
		{
			name: "OK",
			s:    s,
			request: &Todo{
				Message:     "buy some milk",
			},
			mock: func() {
				mock.ExpectPrepare("INSERT INTO todos").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &Todo{
				Id:        1,
				Message: "buy some milk",
			},
			wantErr: false,
		},
		{
			name: "Empty message",
			s: s,
			request: &Todo{
				Message:      "",
			},
			mock: func(){
				mock.ExpectPrepare("INSERT INTO todos").ExpectExec().WillReturnError(errors.New("empty title"))
			},
			wantErr: true,
		},		
		{
			name: "Invalid SQL query",
			s: s,
			request: &Todo{
				Message:     "buy some milk",
			},
			mock: func(){
				//Instead of using todos table, used wrong_table"
				mock.ExpectPrepare("INSERT INTO wrong_table").ExpectExec().WithArgs("message").WillReturnError( errors.New("invalid sql query"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Create(tt.request)
			if (err != nil) != tt.wantErr {
				fmt.Println("this is the error message: ", err.Message())
				t.Errorf("Create() error = %v, wantErr %v, message %v", err, tt.wantErr, err.Message())
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoRepo_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := NewTodoRepository(db)

	tests := []struct {
		name    string
		s       todoRepoInterface
		msgId   int64
		mock    func()
		want    []Todo
		wantErr bool
	}{
		{
			//When everything works as expected
			name:  "OK",
			s:     s,
			mock: func() {
				//We added two rows
				rows := sqlmock.NewRows([]string{"Id", "Message"}).AddRow(1, "buy some milk").AddRow(2, "buy some chocolate")
				mock.ExpectPrepare("SELECT (.+) FROM todos").ExpectQuery().WillReturnRows(rows)
			},
			want: []Todo{
				{
					Id:        1,
					Message:   "buy some milk",
				},
				{
					Id:        2,
					Message:   "buy some chocolate",
				},
			},
		},
		{
			name:  "Invalid SQL Syntax",
			s:     s,
			mock: func() {
				//We added two rows
				_ = sqlmock.NewRows([]string{"Id", "Message"}).AddRow(1, "buy some milk").AddRow(2, "buy some chocolate")
				//"SELECTS" is used instead of "SELECT"
				mock.ExpectPrepare("SELECTS (.+) FROM todos").ExpectQuery().WillReturnError(errors.New("Error when trying to prepare all todos"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}


//When the right number of arguments are passed
//This test is just to improve coverage
func TestTodoRepo_Initialize(t *testing.T) {
	dbdriver :=  "mysql"
	username := "username"
	password := "password"
	host := "host"
	database := "database"
	port := "port"
	dbConnect := TodoRepo.Initialize(dbdriver, username, password, port, host, database)
	fmt.Println("this is the pool: ", dbConnect)
}