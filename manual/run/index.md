---
layout: manual
title: Command "sunplate run"
customHeader: true
links:
  next:
    name: Configuration file for the file watcher / task runner
    url: /manual/run/config.html
---
# Command `sunplate run`
This command starts a file watcher / task runner.
It waits for changes of files in specified directories and executes specified commands
when the changes occur.

This can be useful for:

* (Re)compilation and (re)start of your Go application and its dependencies.
* Compilation of SCSS, LESS, CoffeeScript, concatenation and compression of JS, CSS, etc.

**Just as any other utility of Sunplate toolkit, `sunplate run` is not tied to other components.
It is indepent and you can use it separately from the toolkit.**
