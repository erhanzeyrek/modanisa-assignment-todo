package controllers

import (
	"net/http"
	"strconv"
	"todo-api/domain"
	"todo-api/services"
	"todo-api/utils/error_utils"

	"github.com/gin-gonic/gin"
)

func GetTodoId(msgIdParam string) (int64, error_utils.TodoErr) {
	msgId, msgErr := strconv.ParseInt(msgIdParam, 10, 64)
	if msgErr != nil {
		return 0, error_utils.NewBadRequestError("todo id should be a number")
	}
	return msgId, nil
}

func GetTodo(c *gin.Context) {
	msgId, err := GetTodoId(c.Param("message_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	message, getErr := services.TodoService.GetTodo(msgId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, message)
}

func GetAllTodos(c *gin.Context) {
	messages, getErr := services.TodoService.GetAllTodos()
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, messages)
}

func CreateTodo(c *gin.Context) {
	var message domain.Todo
	if err := c.ShouldBindJSON(&message); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	msg, err := services.TodoService.CreateTodo(&message)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, msg)
}