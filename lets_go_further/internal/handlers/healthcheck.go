package handlers

import (
	"fmt"
	"net/http"
)

func Healthcheck(environment, version string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		fmt.Fprintf(w, `
<div>
	<p><strong>Status:</strong> available</p>
	<p><strong>Environment:</strong> %s</p>
	<p><strong>Version:</strong> %s</p>
</div>
`, environment, version)
	}
}
