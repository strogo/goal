---
layout: manual
title: Deployment
---
> Sunplate does not bring connection / resource management by-default and thus
> all production deployments should have an HTTP proxy that is properly configured
> in front of all Sunplate HTTP requests.



An app that is developed using Sunplate Toolkit does not require any
extra dependencies to be built.
It is a regular Go language project.
And thus, if your are using a version control system to distribute updates,
you may take advantage of existing cartriges, buildpacks, etc.
that bring support of Go.

## Openshift
Use standard [Go language cartridge](https://hub.openshift.com/quickstarts/29-go-language)
when deploying your app to Openshift platform.

## Heroku
Use [Go buildpack](https://devcenter.heroku.com/articles/getting-started-with-go#introduction) for deployment
of your project to Heroku.
