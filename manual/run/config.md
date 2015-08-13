---
layout: manual
title: Configuring file watcher / task runner
category: run
links:
  prev:
    name: General description of file watcher / task runner
    url: /manual/run/index.html
---

## Sections
To specify directories to watch and commands to execute `sunplate.yml`
file at the root of your project is used.
It has the following format:

```yaml
# This section has commands that will be executed at the start of
# `sunplate run` and every time this configuration file is modified.
init: [...]

# This section defines the directories to watch and commands
# to execute when some files there are modified.
watch:
  path/to/dir: [...]
  path/to/another/dir*: [...]
```

As you can see there two main sections. They are:

* `init` contains commands that are executed once per start of `sunplate run` command.
Moreover, they are started again when your `sunplate.yml` file is changed and its reload is required.
* `watch` specifies directories and their lists of commands. Just as you can may expect,
the commands will be executed when some files in those directories are updated.

**Pay attention to the `path/to/another/dir*` path. It has `*` at the end which means
the directory will be watched recursively, i.e. its subdirectories will be watched, too.**

## Commands
For illustration of how real `init` section may look like here is a fragment that
installs `coffeescript` and deletes some temporary directory:

```yaml
init:
  - npm install coffeescript
  - rm -rf ./tmp
```

**All commands will be started from the root of your project.**

As a note, use of commands specific for some platform (such as `rm -rf` above) is discouraged.
It is always better to use cross platform equivalents, if possible.

## Variables
Commands may also include variables. The list of currently supported ones includes:

* `:EXT` - will be replaced by ".exe" when running on Windows and by "" otherwise.

This may be useful when building your application:

```yaml
init:
  - go build -o ./bin/app:EXT ./main.go
```

This will produce "./bin/app.exe" binary when on Windows and "./bin/app" otherwise.

## Builtin commands
Apart from the kinds of commands described above, there is a list of `sunplate run` builtin commands, they are:

* `/start` - start some commands asynchronously
* `/run` - start some commands and wait them to complete
* `/single` - start a new instance of some command asynchronously, make sure the old one is killed
* `/echo` - print something
* `/pass` - do nothing

They can only be used inside `init` and `watch` sections.
There is a demonstration of builtin commands in action:

```yaml
init:
  - /start  list_of_commands # Start commands from the list and don't wait them.
  - /run    list_of_commands # Start commands from list and wait them to complete.
  - /single one_command      # Run exactly one instance of a command from section.

  - /echo "Hello, world"     # Print "Hello, world".
  - /pass                    # Do nothing.

list_of_commands:
  - go build ./

one_command:
  - ./run.sh
```

## Example
A full working example of `sunplate.yml` configuration file is provided below:

```yaml
init:
  - npm install coffeescript # Install CoffeeScript.
  - /run buildApp            # Build application for the first time.
  - /single startApp         # Run a single instance of app.
watch:
  ./controllers*:            # Watch "./controllers" and its subdirectories.
    - /run buildApp          # Rebuild the app on change.
    - /single startApp       # Restart the previously started instance of app.
  ./cscripts:
    # Print a message about recompilation of CS.
    - /echo "CoffeeScript assets will be recompiled..."

    # Recompile CoffeeScript assets.
    - coffeescript ./cscripts --compile --output ./static/jscripts
  ./qwerty:
    - /pass # Do nothing as we haven't decided yet what to do here.
buildApp:
  # Commands for building the project.
  - go build -o ./bin/run:EXT github.com/anonx/sample
startApp: ./bin/run:EXT # A command for start of the app.
```
