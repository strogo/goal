// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/colegion/goal/controllers/templates"

	"github.com/colegion/goal/config"
	"github.com/colegion/goal/strconv"
)

// Templates is an insance of tTemplates that is automatically generated from Templates controller
// being found at "github.com/colegion/goal/controllers/templates/templates.go",
// and contains methods to be used as handler functions.
//
// Templates is a controller that provides support of HTML result
// rendering to your application.
// Use SetTemplatePaths to register templates and
// call c.RenderTemplate from your action to render some.
var Templates tTemplates

// tTemplates is a type with handler methods of Templates controller.
type tTemplates struct {
}

// New allocates (github.com/colegion/goal/controllers/templates).Templates controller,
// then returns it.
func (t tTemplates) New() *contr.Templates {
	c := &contr.Templates{}
	return c
}

// Before calls (github.com/colegion/goal/controllers/templates).Templates.Before.
func (t tTemplates) Before(c *contr.Templates, w http.ResponseWriter, r *http.Request) http.Handler {
	// Call magic Before action of (github.com/colegion/goal/controllers/templates).Templates.
	if res := c.Before( // "Binding" parameters.
	); res != nil {
		return res
	}
	return nil
}

// After is a dump method that always returns nil.
func (t tTemplates) After(c *contr.Templates, w http.ResponseWriter, r *http.Request) http.Handler {
	return nil
}

// Initially is a method that is started by every handler function at the very beginning
// of their execution phase.
func (t tTemplates) Initially(c *contr.Templates, w http.ResponseWriter, r *http.Request, a []string) (finish bool) {
	// Call magic Initially method of (github.com/colegion/goal/controllers/templates).Templates.
	return c.Initially(w, r, a)
}

// Finally is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tTemplates) Finally(c *contr.Templates, w http.ResponseWriter, r *http.Request, a []string) (finish bool) {
	return
}

// Render is a handler that was generated automatically.
// It calls Before, After, Finally methods, and Render action found at
// github.com/colegion/goal/controllers/templates/templates.go
// in appropriate order.
//
// Render is an equivalent of
// RenderTemplate(ControllerName+"/"+ActionName+".html").
func (t tTemplates) Render(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Templates.New()
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	a := []string{"Templates", "Render"}
	defer Templates.Finally(c, w, r, a)
	if finish := Templates.Initially(c, w, r, a); finish {
		return
	}
	if res := Templates.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.Render( // "Binding" parameters.
	); res != nil {
		h = res
		return
	}
	if res := Templates.After(c, w, r); res != nil {
		h = res
	}
}

// RenderTemplate is a handler that was generated automatically.
// It calls Before, After, Finally methods, and RenderTemplate action found at
// github.com/colegion/goal/controllers/templates/templates.go
// in appropriate order.
//
// RenderTemplate is an action that gets a path to template
// and renders it using data from Context.
func (t tTemplates) RenderTemplate(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Templates.New()
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	a := []string{"Templates", "RenderTemplate"}
	defer Templates.Finally(c, w, r, a)
	if finish := Templates.Initially(c, w, r, a); finish {
		return
	}
	if res := Templates.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.RenderTemplate( // "Binding" parameters.
		strconv.String(r.Form, "templatePath"),
	); res != nil {
		h = res
		return
	}
	if res := Templates.After(c, w, r); res != nil {
		h = res
	}
}

// RenderError is a handler that was generated automatically.
// It calls Before, After, Finally methods, and RenderError action found at
// github.com/colegion/goal/controllers/templates/templates.go
// in appropriate order.
//
// RenderError is an action that renders Error 500 page.
func (t tTemplates) RenderError(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Templates.New()
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	a := []string{"Templates", "RenderError"}
	defer Templates.Finally(c, w, r, a)
	if finish := Templates.Initially(c, w, r, a); finish {
		return
	}
	if res := Templates.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.RenderError( // "Binding" parameters.
	); res != nil {
		h = res
		return
	}
	if res := Templates.After(c, w, r); res != nil {
		h = res
	}
}

// RenderNotFound is a handler that was generated automatically.
// It calls Before, After, Finally methods, and RenderNotFound action found at
// github.com/colegion/goal/controllers/templates/templates.go
// in appropriate order.
//
// RenderNotFound is an action that renders Error 404 page.
func (t tTemplates) RenderNotFound(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Templates.New()
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	a := []string{"Templates", "RenderNotFound"}
	defer Templates.Finally(c, w, r, a)
	if finish := Templates.Initially(c, w, r, a); finish {
		return
	}
	if res := Templates.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.RenderNotFound( // "Binding" parameters.
	); res != nil {
		h = res
		return
	}
	if res := Templates.After(c, w, r); res != nil {
		h = res
	}
}

// Init is used to initialize controllers of "github.com/colegion/goal/controllers/templates"
// and its parents.
func Init(g config.Getter) {
	initTemplates(g)
	contr.Init(g)
}

func initTemplates(g config.Getter) {
}

func init() {
	_ = strconv.MeaningOfLife
}
