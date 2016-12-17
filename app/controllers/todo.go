package controllers

import (
	"github.com/revel/revel"
	"github.com/spenserpothier/revel_testapp/models"
	//"github.com/spenserpothier/revel_testapp/app/routes"
	"strconv"
	"github.com/spenserpothier/revel_testapp/app/routes"
)

type Todo struct {
	*revel.Controller
}

func (c Todo) Index() revel.Result {
	return c.Render(todoList)
}

var todoList = []models.Todo{}

func (c Todo) AddToDo(todo models.Todo) revel.Result {
	todo.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Todo.Index)
	}

	todo.Number =len(todoList)
	todo.Number++
	todo.Creator = c.Session["userName"]
	todoList = append(todoList, todo)

	return c.Redirect(Todo.Index)
}

func (c Todo) CompleteTodo() revel.Result {
	c.Request.ParseForm()
	for i, _ := range c.Request.Form{
		t, _ := strconv.Atoi(i)
		t--
		todoList[t].Completed = true
	}

	return c.Redirect(Todo.Index)
}

func (c Todo) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Index())
}

func (c Todo) CheckLogin() revel.Result {
	if _, ok := c.Session["userName"]; ok {
		return nil
	}

	c.Flash.Error("Please Login First!")
	return c.Redirect(routes.App.Index())
}