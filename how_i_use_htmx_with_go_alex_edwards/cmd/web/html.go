package main

import (
	"bytes"
	"html/template"
	"io/fs"
	"net/http"
	"time"
)

type htmlRenderer struct {
	templateFS      fs.FS
	sharedTemplates *template.Template
}

// The newHTMLRenderer function creates a new htmlRenderer containing a shared
// set of parsed templates with support for any custom template functions
func newHTMLRenderer(templateFS fs.FS, sharedTemplateFiles ...string) (*htmlRenderer, error) {
	funcs := template.FuncMap{
		"now": time.Now,
		// Other custom template functions go here...
	}
	sharedTemplates, err := template.New("").Funcs(funcs).ParseFS(templateFS, sharedTemplateFiles...)
	if err != nil {
		return nil, err
	}

	r := &htmlRenderer{
		templateFS:      templateFS,
		sharedTemplates: sharedTemplates,
	}
	return r, nil
}

// The render method clones the shared template set, optionally parses additional
// templates, executes the named template with the supplied data, and writes the
// response.
func (h *htmlRenderer) render(w http.ResponseWriter, status int, data any, templateName string, additionalTemplateFiles ...string) error {
	// Clone the startup templates so request-specific parsing does not mutate the shared template set.
	ts, err := h.sharedTemplates.Clone()
	if err != nil {
		return err
	}

	// Add templates needed only by this render
	if len(additionalTemplateFiles) > 0 {
		ts, err = ts.ParseFS(h.templateFS, additionalTemplateFiles...)
		if err != nil {
			return err
		}
	}

	buf := new(bytes.Buffer)

	err = ts.ExecuteTemplate(buf, templateName, data)
	if err != nil {
		return err
	}
	// We could set the Vary: HX-Request header in searchUsers, but I think it's easier to just always set it on all responses in the render() function. It does mean that we'll be setting the Vary header on all responses — including those from our home handler and listUsers — which isn't strictly necessary and a little bit wasteful. But I think it's worth it to avoid having to remember setting the Vary header correctly in individual handlers, and the risk of bugs that forgetting it may cause.

	w.Header().Add("Vary", "HX-Request")
	w.WriteHeader(status)
	buf.WriteTo(w)

	return nil
}
