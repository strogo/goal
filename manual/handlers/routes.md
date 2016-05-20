---
layout: manual
title: Routes
category: handlers
---
Optionally, actions may have routes associated with them.
For illustration:

```go
//@get /greet/:name
func (c *Application) Greet(name string) http.Handler {
	...
}
```
The route must be placed in the comment above the appropriate action.
The list of supported methods includes:

* GET, HEAD, POST, PUT, DELETE, TRACE, OPTIONS, CONNECT, PATCH

Apart from that there is also special ROUTE method that is equivalent of
saying "all routes":

```go
//@route /
func (c *Application) Index() http.Handler {
	...
}
```

Actions are allowed to have any number of routes:

```go
//@get /
//@post /helloworld
//@delete /something/else
func (c *Application) HelloWorld() http.Handler {
	...
}
```

## Inheritance
When importing actions from some third party controller, it is possible to decide
how you want to mount them to your application.
E.g. let's assume there is a third party controller that has the following actions and routes:

```go
//@get /signin
func (c *Accounts) SignIn() http.Handler {
	...
}
```
To make this part of your application and avoid collision you can do the following:

```go
type MyApp struct {
	*Accounts `@route:"/accounts"`
}
```
This way, `Account`'s action will be accessible in your application as `/accounts/signin`.
Instead of `@route` struct tag you could also use any other supported type.
That would mean that you want to mount routes of that specific types.
