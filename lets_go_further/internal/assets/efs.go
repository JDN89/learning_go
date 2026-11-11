package assets

import (
	"embed"
	"io/fs"
)

// NOTE: allows embeding files in the go binary
// here we embed the contents of the 'html' and 'static' directories in the compiled binary
// holds an in-memory filesystem containing those embedded files.

//go:embed "html" "static"
var files embed.FS

// NOTE: It provides a clear separation between the static and HTML files when we are using them from our Go code. Code that is intended to only work with our static files won't have unnecessary access to our HTML files, and vice-versa.
// Code using the HTMLFiles and StaticFiles filesystems doesn't need to include the html/ or static/ path prefix when opening files.

var (
	HTMLFiles   = sub(files, "html")
	StaticFiles = sub(files, "static")
)

// sub accepts the embeded filesystem adn the directory to access
func sub(f embed.FS, dir string) fs.FS {
	// create a filesystem rooted at the specified directory
	sub, err := fs.Sub(f, dir)
	if err != nil {
		panic(err)
	}
	return sub
}
