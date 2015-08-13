---
layout: manual
title: Introduction to Sunplate
---
Sunplate is a set of independent utilities that may be used together or separately.
All of them form a full stack MVC (Model - View - [Controller](handlers/controllers.html)) web framework.

The list of currently available utilities includes:

* `sunplate create` creates a new skeleton application with opinionated [structure](create/index.html#layout) and parameters.
This is a command for easy and fast start of a new project.
* `sunplate run` runs a file watcher / task runner that is responsible for (re)compilation / (re)start  of your project and,
if neccessary, (re)build of client-side assets. [Configuration file](run/index.html#sunplate.yml) is used to describe how
to compile, build, start things.
* `sunplate generate handlers` generates a package with HTTP handler functions from your
[controllers](handlers/controllers.html). The generated package is fully compatable with the standard library
and thus can be used with any router.
