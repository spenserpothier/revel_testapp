package controllers

import (
	"github.com/revel/revel"
	"github.com/spenserpothier/revel_testapp/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Login(user *models.User) revel.Result {
	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	c.Session["userName"] = user.Name
	c.Session.SetNoExpiration()

	c.Flash.Success("Welcome, " + user.Name + "!")

	return c.Redirect(Todo.Index)
}
