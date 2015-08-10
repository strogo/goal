---
layout: main
title: High productivity web development in Go language - Sunplate Toolkit
---
## Overview
### The idea
Sunplate is a set of independent utilities. 

* `sunplate create` - create a skeleton application.
* `sunplate run` - start a file watcher and task runner.
* `sunplate generate` - automatic generation of Go code (used together with `go generate`).

### Controllers
> Command `sunplate generate handlers`
> generates a `handlers` package from your controllers.
> The generated package will contain Go handler functions
> that can be used as `handlers.ControllerName.ActionName`.
>
> Read below about how your controllers should look like.

Controller is any structure that has [actions](#actions). Example:

```go
// Profiles is a sample controller.
type Profiles struct {
	*Controller
}
```

An inheritance of controllers is possible. Just as illustrated above `Profiles`
has a parent `Controller`.

### Actions
Action is a controller's exported method that returns a type implementing
`github.com/anonx/sunplate/action.Result` as its first result.

```go
import "github.com/anonx/sunplate/action"

// Index is a sample action of Profiles controller.
func (c *Profiles) Index() action.Result {
	...
}
```

Actions are allowed to return any number of arguments. The only requirement is the first
one must be `action.Result`.
Moreover, they may get any number of builtin type arguments. They will be extracted
from a URL and/or `request.Form`.

```go
// History is a sample of action that gets two arguments
// and may return an error as a second result.
func (c *Profiles) History(username string, page int) (action.Result, error) {
	...
}
```

### Magic actions
There are two actions with special meaning. They are
[`Before`](#magic-before-action) and [`After`](#magic-after-action).
If presented, they will be automatically executed with every request.

#### Magic Before action
Before action will be started before an [action](#action). If `Before` returns a non-nil result,
[action](#action) or [magic After action](#magic-after-action) will not be started.

```go
// Before is a magic method that is automatically started before
// any other action of Profiles controller.
func (c *Profiles) Before() action.Result {
	if c.NotAuthorize() {
		// Login form will be rendered.
		return c.RenderTemplate("profiles/login.html")
	}
	// Go on executing an action and magic After action.
	return nil
}
```

#### Magic After action
After is equivalent to [magic Before action](#magic-before-action) except
it is started after an [action](#action).
And if action returns a non nil result, After action will not be started.


### Magic methods
Besides magic actions there are also two magic methods.

#### Magic Initially method
`Initially` is started before any [action](#action) with every request. It gets two arguments:
`http.ResponseWriter` and `*http.Request`. And returns a single result: `finish bool`.
If `finish` is `true` that means no other magic `Initially` methods or actions
shall be started.

```go
// Initially is a sample magic method.
func (c *Profiles) Initially(w http.ResponseWriter, *r.Request) (finish bool) {
	return false // Go on executing other methods and actions.
}
```

#### Magic Finally method
`Finally` is equivalent of [Initially](#magic-initially-method) except it is started at the very end.
And it will be executed even if some of the actions or previous magic methods
panic. If `Finally` returns `true` that means no other magic `Finally` methods are
expected to be started.

```go
// Finally is a sample magic method that catches panics.
func (c *Profiles) Finally(w http.ResponseWriter, *r.Request) bool {
	if err := recover(); err != nil {
		...
		return true
	}
	return false
}
```

### Inheritance
There is a parent controller with magic actions / methods:

```go
type ParentController struct {
}

func (c *ParentController) Before() action.Result {
	return nil
}

func (c *ParentController) Finally(http.ResponseWriter, *http.Request) bool {
	return false
}
```

and our child controller that embeds the parent has some magic actions, too:

```go
type ChildController struct {
	*ParentController
}

func (c *ChildController) Before() action.Result {
	return nil
}

func (c *ChildController) Index() action.Result {
	return nil
}
```

So, as a result methods will be called in the following order
when trying to access `Index`:

1. `ParentController.Before`
2. `ChildController.Before` (if 1 returned `nil`)
3. `ChildController.Index` (if 1 and 2 returned `nil`)
4. `ParentController.Finally`

### Task runner / file watcher

> Command `sunplate run` starts a file watcher / task runner that
> reads `sunplate.yml` file at the root of your project
> and uses it to (re)build / (re)start your project and
> necessary dependencies and CSS/CoffeeScript/Etc. assets.
>
> The format of `sunplate.yml` configuration file is described
> below.

```yaml
# Commands that are expected to be started at the beginning.
# You may use it to install / update some third party build components.
init:
	- npm install coffeescript

	# Builtin commands that may be used inside "init" or "watch" section.
	- /pass                    # Do nothing
	- /echo   "Hello, world"   # Print something
	- /start  list_of_commands # Start commands asynchronously.
	- /run    list_of_commands # Run commands and wait them to complete.
	- /single single_command   # Start a single instance of a command, make sure the previous one is killed.

# Watch the directories and execute the command every time
# some file has been modified there.
watch:
	path/to/dir:
		- coffeescript smth
	path/to/another/dir*: # Asterisk at the end means scan recursively.
		- go build smth

# A sample section with a list of commands.
list_of_commands:
	- command_one
	- command_two

# A sample section with just one command.
single_command: some_command
```
