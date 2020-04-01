package main

import (
	"net/http"
)

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	// standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	// dynamicMiddleware := alice.New(app.session.Enable, noSurf)

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest)
	dynamicMiddleware := alice.New()

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
