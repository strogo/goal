package controllers

import (
	"github.com/anonx/sunplate/action"

	"github.com/revel/revel/testing"
)

// App is a sample controller.
type App struct {
	*Controller
	*NotController
	*NotController1
}

// NotController is not a controller as it doesn't have methods.
type NotController struct {
}

// NotController1 does have methods but it is still not a controller
// as there are not actions among those methods.
type NotController1 struct {
}

func (c NotController1) test() {
}

// HelloWorld is a sample action.
func (c App) HelloWorld(page int) action.Result {
	return nil
}

// Finally is a magic method that is executed after every request.
func (c *Controller) Finally(name string) action.Result {
	return nil
}

// UnsupportedAction is not an action as it requires argument that is not
// of builtin type.
func (c Controller) UnsupportedAction(t testing.TestSuite) action.Result {
	return nil
}

// Smth is not an action as it returns nothing.
func (c App) Smth() {
}

func init() {
}