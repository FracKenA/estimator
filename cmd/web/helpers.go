package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

import (
	// "github.com/justinas/nosurf"
	"github.com/sirupsen/logrus"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.log.WithFields(logrus.Fields{
		"event": "serverError",
	}).Error(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) addDefaultData(td *templateData, _ *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}

	// td.CSRFToken = nosurf.Token(r)
	td.CurrentYear = time.Now().Year()
	// td.Flash = app.session.PopString(r, "flash")
	return td
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exist", name))
		return
	}
	buf := new(bytes.Buffer)
	err := ts.Execute(buf, app.addDefaultData(td, r))

	if err != nil {
		app.serverError(w, err)
		return
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		app.serverError(w, err)
		return
	}

}

/*
NOTE kept in case auth needed later
func (app *application) isAuthenticated(r *http.Request) bool {
	isAuthenticated := false
	isAuthenticated, ok := r.Context().Value(contextKeyIsAuthenticated).(bool)

	if !ok {
		return false
	}
	return isAuthenticated
}

func (app *application) isAdmin(r *http.Request) bool {
	isAdmin := false
	isAdmin, ok := r.Context().Value(contextKeyIsAdmin).(bool)

	if !ok {
		return false
	}
	return isAdmin
}
*/
