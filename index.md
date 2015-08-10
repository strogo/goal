---
layout: main
title: High productivity web development in Go language - Sunplate Toolkit
---
## Overview
### The idea
Sunplate is a set of independent utilities. The list of currently supported ones includes:

* `sunplate create {path}` - create a skeleton application in the specified destination.
* `sunplate run {path}` - start a file watcher and task runner.
It opens `sunplate.yml` configuration at the root of your project to find out what to watch
and how to (re)start / (re)build your application. Format of the file is described below.
* `sunplate generate handlers` - generate standard Go handler functions from your controllers.
* `sunplate generate views` - generate a list of found templates.

### Concept
