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

> Command `sunplate generate handlers`
> generates a `handlers` package from your controllers.
> The generated package will contain Go handler functions
> that can be used as `handlers.ControllerName.ActionName`.
>
> Read below about how your controllers should look like.

### Controllers
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
[`Before`](#before-magic-action) and [`After`](#after-magic-action).
If presented, they will be automatically executed with every request.

#### Before magic action
Before action will be started before an [action](#action). If `Before` returns a non-nil result,
[action](#action) or [`After` magic action](#after-magic-action) will not be started.

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

#### After magic action
After is equivalent to [Before magic action](#before-magic-action) except
it is started after an [action](#action).
And if action returns a non nil result, After action will not be started.


### Magic methods
Besides magic actions there are also two magic methods.

#### Initially magic method
Initially is started before any [action](#action) with every request. It gets two arguments:
`http.ResponseWriter` and `*http.Request`. And returns a single result: `finish bool`.
If `finish` is `true` that means no other Initially magic methods or actions
shall be started.

```go
// Initially is a sample magic method.
func (c *Profiles) Initially(w http.ResponseWriter, *r.Request) (finish bool) {
	return false // Go on executing other methods and actions.
}
```

#### Finally magic method
Finally is equivalent of [Initially](#initially-magic-action) except it is started at the very end.
And it will be executed even if some of the actions or previous magic methods
panic. If Finally returns `true` that means no other magic Finally methods are
expected to be started.

```go
// Finally is a sample magic method that catches panics.
func (c *Profiles) Initially(w http.ResponseWriter, *r.Request) bool {
	if err := recover(); err != nil {
		...
		return true
	}
	return false
}
```
