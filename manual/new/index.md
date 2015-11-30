---
layout: manual
title: Command "goal new"
customHeader: true
---
# Command `goal new`
This command is the way to start development easily and fast. It generates a new skeleton application
with [opinionated structure](#default-layout) and parameters and saves it to the requested destination.

Example of its use is:

```bash
goal new github.com/colegion/sample
```

This will make a `sample/` directory inside `$GOPATH/src/github.com/colegion` and create all neccessary files.
Alternatively, you can use relative path:

```bash
goal new ./colegion/sample
```

**Please, note that during development your project must be
located within [`$GOPATH/src`](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable) directory.**

## Default Layout
A project generated by `goal new` will have the following structure:

* `assets/` - automatically generated assets
  * `handlers/` - package and its subpackages generated by [`goal generate handlers`](../handlers/)
  * `views/` - package generated by [`goal generate views`](../views/)
* `controllers/` - [controllers](../handlers/controllers.html) of your application
  * `app.go` - sample controller
  * `controllers.go` - registration of [builtin controllers](../controllers/) and their initialization
  * `init.go`
* `main.go` - entry point of your application: initialization of components, creation and start of a server
* `nix.go` - code for graceful restarts / shutdowns of your app using
[gracehttp](https://github.com/facebookgo/grace)
* `routes/` - package for mointing handler functions to HTTP Method / Path values
* `static/` - static public assets
* `goal.yml` - configuration file of [`goal run`](../run/) file watcher / task runner
* `views/` - server-side templates of your application
* `win.go` - fallback code for Windows that does not support graceful restarts / shutdowns

**The layout is recommended but not required.**
If you feel like a different structure would fit your project better apply it.
All you will have to do is to make sure `//go:generate` directives (by-default are located inside
`controllers/init.go` and `controllers/controllers.go`) and `goal.yml` reflect your changes.