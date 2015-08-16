---
layout: manual
title: Actions
category: handlers
links:
  prev:
    name: Everything about controllers
    url: /manual/handlers/controllers.html
---
Action is a [controller](controllers.html)'s exported method that returns an
`http.Handler` (of `net/http`) as its first result.

```go
import "net/http"

// Index is a sample actions of Profiles controller.
func (c *Profiles) Index() http.Handler {
	...
}
```

Actions are allowed to return any number of arguments. The only requirement is the first
one must always be `http.Handler`.
Moreover, they may get any number of arguments of the following types:

* `bool`
* `string`
* `int`, `int8`, `int16`, `int32`, `int64`
* `uint`, `uint8`, `uint16`, `uint32`, `uint64`
* `float32`, `float64`
* `[]bool`
* `[]string`
* `[]int`, `[]int8`, `[]int16`, `[]int32`, `[]int64`
* `[]uint`, `[]uint8`, `[]uint16`, `[]uint32`, `[]uint64`
* `[]float32`, `[]float64`

The arguments will be initialized with the values extracted from the following sources (in that order):

* The URL `/:path` parameters
* The URL `?query` parameters
* Submitted `Form` values

```go
// History is a sample of action that gets two arguments
// and may return an error as a second result.
func (c *Profiles) History(username string, page int) (http.Handler, error) {
	...
}
```

## Magic actions
Magic actions are actions with special meaning. If presented, they are invocated automatically.
If action / magic action returns a `non-nil` result, no subsequent actions will be started, read more
about it in the description of [request life cycle](controllers.html#request-life-cycle).

### `Before`
`Before` magic method is invocated before [action](#actions) invocation.

```go
// Before is a magic method that is automatically started before
// any other action of Profiles controller.
func (c *Profiles) Before() http.Handler {
	if c.NotAuthorize() {
		// Login form will be rendered.
		return c.RenderTemplate("profiles/login.html")
	}
	// Go on executing an action and magic After action.
	return nil
}
```

### `After`
`After` is equivalent to `Before` except it is invocated after [action](#actions) invocation.
