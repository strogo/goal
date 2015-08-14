---
layout: manual
title: Command "sunplate generate handlers"
customHeader: true
---
# Command `sunplate generate handlers`
This is the command that scans your [controllers](controllers.html) package and generates
a package with handler functions where there is a handler function per every
[action](actions.html) of your controllers.

So, if your have the following controller:

```go
// Application is a controller.
type Application struct {
}

// Index is an action.
func (c *Application) Index() action.Result {
	...
}
```

You would get the `handlers` package with `Index(w http.ResponseWriter, r *http.Request)`
that can be accessed as `handlers.Application.Index`.

## Supported arguments

* `--input` - a path where controllers can be found. Default is `./controllers`.
* `--output` - a path where to save generated package files. Default is `./assets/handlers`.
**Be careful, all files in this directory will be removed before generation of `handlers` package.**
* `--package` - name of the package to generate. Default is `handlers`.

**Important: All paths are expected to be relative to the root of your project.**
