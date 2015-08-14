---
layout: manual
title: Standard controllers
category: controllers
---
Standard controllers are the controllers that are included in your
[automatically generated](new) app by-default. Though, you can simply remove / replace
any of them from your app if your like.

## Template
Template is template system that is based on Go's standard `html/template` package.
It is however has a difference:

* Every template must either have `base` block inside or have `Base.html` file with `base` block
in the same directory, or in one of the upper directories.

For illustration, there is a sample layout and templates:

* `Base.html`
* `Profiles/Index.html`

```html
{% raw %}
<!-- Sample of Base.html -->
{% define "base" %}
<!DOCTYPE html>
<html>
<head>
	<title>{% template "title" . %}</title>
</head>
<body>
	{% template "body" . %}
</body>
</html>
{% end %}
{% endraw %}
```

```html
{% raw %}
<!-- Sample of Profiles/Index.html -->
{% define "title" %}Some title{% end %}
{% define "body" %}
	Hello, world!
{% end %}
{% endraw %}
```

Template controller will bring the following methods to your actions:

* `RenderTemplate("Path/To/File.html")`

## Params
Params controller parses request and makes it available as `Request` field in your controller.
That let you access parameters as `c.Request.FormValue("somekey")`.
