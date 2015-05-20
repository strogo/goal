package template

import (
	"net/http"
)

// HTML is a result that is returned from actions by default.
// This struct implements action.Result interface.
type HTML struct {
}

// Apply writes to response the result received from action.
func (t *HTML) Apply(w http.ResponseWriter, r *http.Request) {
}

// Finish means that ResponseWriter should be used to return
// result to the user immidiately.
func (t *HTML) Finish() bool {
	return true
}