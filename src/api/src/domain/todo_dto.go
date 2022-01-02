package domain

import (
	"strings"
	"todo-api/utils/error_utils"
)

type Todo struct {
	Id        int64     `json:"id"`
	Message     string    `json:"message"`
}

func (m *Todo) Validate() error_utils.TodoErr {
	m.Message = strings.TrimSpace(m.Message)
	if m.Message == "" {
		return error_utils.NewUnprocessibleEntityError("Please enter a valid todo message")
	}
	return nil
}