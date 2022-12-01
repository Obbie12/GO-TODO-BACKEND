package handlers

import (
	"go-todo-backend/models"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllTodosHandler(c echo.Context) error {
	var todo []models.Todo
	err := models.GetAllTodos(&todo)
	if err != nil {
		return c.String(http.StatusNotFound, "Todo note was not found")
	} else {
		return c.JSON(http.StatusOK, todo)
	}
}

func CreateTodoHandler(c echo.Context) (err error) {
	var todo models.Todo
	c.Bind(&todo)
	err = models.CreateATodo(&todo)
	if err != nil {
		return c.String(http.StatusNotFound, "Todo note was not found")
	} else {
		return c.JSON(http.StatusCreated, todo)
	}
}

func GetTodoHandler(c echo.Context) (err error) {
	id, err := GetTodoId(c)

	if err != nil {
		return err
	}

	var todo models.Todo
	err = models.GetATodo(&todo, id)
	if err != nil {
		return c.String(http.StatusNotFound, "Todo note was not found")
	} else {
		return c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodoHandler(c echo.Context) (err error) {
	var todo models.Todo
	id, err := GetTodoId(c)

	if err != nil {
		return err
	}

	err = models.DeleteATodo(&todo, id)
	if err != nil {
		return c.String(http.StatusNotFound, "Todo note was not found")
	} else {
		return c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodoHandler(c echo.Context) (err error) {
	var todo models.Todo
	id, err := GetTodoId(c)

	if err != nil {
		return err
	}

	err = models.GetATodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	c.Bind(&todo)

	err = models.UpdateATodo(&todo, id)
	if err != nil {
		return c.String(http.StatusNotFound, "Todo note was not found")
	} else {
		return c.JSON(http.StatusOK, todo)
	}
}

func GetTodoId(c echo.Context) (id int, err error) {
	rawId := c.Param("id")
	id, err = strconv.Atoi(rawId)
	return
}
