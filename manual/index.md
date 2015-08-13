---
layout: manual
title: Introduction to Sunplate
---
Sunplate is a set of independent utilities that may be used together or separately.
All of them form a full stack MVC (Model - View - [Controller](handlers/controllers.html)) web framework.

The list of currently available utilities includes:

* `sunplate new` creates a new skeleton application with opinionated [structure](new/index.html#layout) and parameters.
This is a command for easy and fast start of a new project.
* `sunplate run` runs a file watcher / task runner that is responsible for (re)compilation / (re)start  of your project and,
if neccessary, (re)build of client-side assets. [Configuration file](run/index.html#sunplate.yml) is used to describe how
to compile, build, start things.
* `sunplate generate handlers` generates a package with HTTP handler functions from your
[controllers](handlers/controllers.html). The generated package is fully compatable with the standard library
and thus can be used with any router.

## Getting started
To get started with Sunplate, first install it:

```bash
go get -u github.com/anonx/sunplate
```

After you have the Toolkit installed, create a new project.
`path/to/your/project` must not exist at the time of running this command:

```bash
sunplate new path/to/your/project
```

**Please, note that both Sunplate toolkit and your project (during development) must be
located within [`$GOPATH/src`](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable) directory.**


To start the project you have just created use:

```bash
sunplate run path/to/your/project
```

Gongratulations, your skeleton app is ready and running. Start making changes.
Open [`http://127.0.0.1:8080`](http://127.0.0.1:8080) to view results.
