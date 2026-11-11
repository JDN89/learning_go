package main

import (
	"log/slog"
	"net/http"

	"github.com/JDN89/learning_go/internal/server"
)

// TODO move to main.go
type Application struct {
	config config
	logger *slog.Logger
	html   *server.HTMLRenderer
}

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	// Render the base template and append home.tmpl
	err := app.html.Render(w, http.StatusOK, nil, "base", "pages/home.tmpl")
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

// func (app *application) gopher(w http.ResponseWriter, r *http.Request) {
// 	// Render the base template and append home.tmpl
// 	width := 100
// 	err := app.html.render(w, http.StatusOK, width, "partial:image:gopher")
// 	if err != nil {
// 		app.logger.Error(err.Error())
// 		http.Error(w, http.StatusText(500), 500)
// 	}
// }
//
// // Define a user type. The fields need to be exported so that we can reference
// // them in our HTML templates.
// type user struct {
// 	Name     string
// 	Email    string
// 	IsGopher bool
// }
//
// // Create a hardcoded list of users.
// var users = []user{
// 	{"Alice Madsen", "alice.madsen@example.com", true},
// 	{"Theo Thatcher", "theo.thatcher@example.com", true},
// 	{"Maxwell Albright", "maxwell.albright@example.com", false},
// 	{"Ruby Thompson", "ruby.thompson@example.com", false},
// 	{"Leona Rowan", "leona.rowan@example.com", false},
// 	{"Alicia Lennox", "alicia.lennox@example.com", true},
// 	{"Ruben Mason", "ruben.mason@example.com", false},
// 	{"Leo Reynolds", "leo.reynolds@example.com", false},
// 	{"Max Lester", "max.lester@example.com", true},
// 	{"Theodore Allister", "theodore.allister@example.com", false},
// }
//
// func (app *application) listUsers(w http.ResponseWriter, r *http.Request) {
// 	// Render a full HTML page containing the content from "pages/users.tmpl"
// 	// and all user details.
// 	err := app.html.render(w, 200, users, "base", "pages/users.tmpl")
// 	if err != nil {
// 		app.logger.Error(err.Error())
// 		http.Error(w, http.StatusText(500), 500)
// 	}
// }

func isHTMXRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}

//NOTE two ways of redirecting check https://www.alexedwards.net/blog/how-i-use-htmx-with-go
// func redirect(w http.ResponseWriter, r *http.Request, url string, code int) {
// 	if isHTMXRequest(r) {
// 		w.Header().Set("HX-Redirect", url)
// 		w.WriteHeader(http.StatusNoContent)
// 		return
// 	}
//
// 	http.Redirect(w, r, url, code)
// }
// func redirect(w http.ResponseWriter, r *http.Request, url string, code int) {
//     if isHTMXRequest(r) {
//         w.Header().Set("HX-Location", url)
//         w.WriteHeader(http.StatusNoContent)
//         return
//     }
//
//     http.Redirect(w, r, url, code)
// }

// func (app *application) searchUsers(w http.ResponseWriter, r *http.Request) {
// 	query := r.FormValue("query")
//
// 	var matches []user
//
// 	if query == "" {
// 		matches = users
// 	} else {
// 		for _, u := range users {
// 			if strings.Contains(u.Name, query) || strings.Contains(u.Email, query) {
// 				matches = append(matches, u)
// 			}
// 		}
// 	}
//
// 	// Render the base template by default.
// 	template := "base"
//
// 	// But if the request is coming from HTMX, render the users:rows template instead.
// 	if isHTMXRequest(r) {
// 		template = "partial:users:rows"
// 	}
//
// 	err := app.html.render(w, 200, matches, template, "pages/users.tmpl")
// 	if err != nil {
// 		app.logger.Error(err.Error())
// 		http.Error(w, http.StatusText(500), 500)
// 	}
// }

// TODO

// Note: I've been deliberately keeping this web application simple so that the focus is on templates and templating. In a real application, you might want to merge listUsers and searchUsers into a single handler, create some centralized helpers for error handling and logging, use middleware to add Content Security Policy headers and recover panics, set appropriate server timeouts, etc.
