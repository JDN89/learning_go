package main

import (
	"encoding/json/v2"
	"net/http"
)

// NOTE in go methods can be defined in different file from where the Type is defined as long as they are in the same package
func (app *Application) Healthcheck(environment, version string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"status":      "available",
			"environment": environment,
			"version":     version,
		}

		js, err := json.Marshal(data)
		if err != nil {
			app.logger.Error(err.Error())
			http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
			return
		}

		js = append(js, '\n')

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, _ = w.Write(js)

		// The previous HTML response:
		// w.Header().Set("Content-Type", "text/html; charset=utf-8")
		//
		// fmt.Fprintf(w, `
		// <div>
		// 	<p><strong>Status:</strong> available</p>
		// 	<p><strong>Environment:</strong> %s</p>
		// 	<p><strong>Version:</strong> %s</p>
		// </div>
		// `, environment, version)
	}
}
