package main

import (
	"net/http"
	"net/url"
)

// NOTE helper to see the current browserURL, that the user is seeing.
// Because HTMX makes AJAX requests and swaps in responses without changing the browser URL (unless you use the hx-boost, hx-push-url or hx-replace-url attributes), the request URL that we see in our Go handlers when accessing r.URL may be different to the one that the user is seeing in their browser.
func browserURL(r *http.Request) (*url.URL, error) {
	cu := r.Header.Get("HX-Current-URL")
	if cu != "" {
		return url.Parse(cu)
	}

	return r.URL, nil
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
