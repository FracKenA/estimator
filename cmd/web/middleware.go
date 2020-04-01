package main

import (
	"fmt"
	"net/http"
)

import(
	//"github.com/justinas/nosurf" Used in unused function
	"github.com/sirupsen/logrus"
)

/*
NOTE: Unused functions.
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")
		next.ServeHTTP(w, r)
	})
}

func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})
	return csrfHandler
}

 */

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.log.WithFields(logrus.Fields{
			"event":      "request",
			"remoteaddr": r.RemoteAddr,
			"method":     r.Method,
			"URI":        r.URL.RequestURI(),
			"protocol":   r.Proto,
		}).Info("Request")
		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

/*
NOTE kept in case authentication is needed later

func (app *application) requireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.isAuthenticated(r) {
			http.Redirect(w, r, "/user/login", 302)
			return
		}
		w.Header().Add("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.isAdmin(r) {
			// TODO change this redirect and add flash
			http.Redirect(w, r, "/", 302)
			return
		}
		w.Header().Add("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			exists := app.session.Exists(r, "authenticatedUserID")
			if !exists {
				next.ServeHTTP(w, r)
				return
			}
			user, err := app.users.Get(app.session.GetInt(r, "authenticatedUserID"))
			if errors.Is(err, models.ErrNoRecord) || !user.Active {
				app.session.Remove(r, "authenticatedUserID")
				next.ServeHTTP(w, r)
				return
			} else if err != nil {
				app.serverError(w, err)
				return
			}
			ctx := context.WithValue(r.Context(), contextKeyIsAuthenticated, true)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
}
*/
