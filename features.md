---
layout: default
title: Features
---

Sunplate, being mostly inspired by Revel Framework and its discussions,
is built around the concept of controllers and actions. However, as opposed to
Revel and other high level frameworks, Sunplate:

#### DOES NOT
* use runtime reflection
* require your app to import monolithic dependencies

Instead, Sunplate uses code generation and `go generate` mechanism.
That let us achieve:

  * type safety
  * minimalism of dependencies
  * compatability with the standard library
  * productivity for the end-developers

Sunplate is very customizable (it is possible to bring your own router, template system, etc.),
but without prejudice to the easiness and seamless of experience thanks to good defaults
of skeleton application.
