---
layout: default
title: Features
---

Sunplate, being mostly inspired by Revel Framework and its discussions,
is built around the concept of controllers and actions. However, as opposed to
Revel and other high level frameworks Sunplate:

#### DOES NOT
* use runtime reflection
* require your app to import monolithic dependencies

#### DOES
* use `go generate` mechanism to achieve:
  * type safety
  * minimalism of dependencies
  * compatability with the standard library
  * productivity for the end-developers
