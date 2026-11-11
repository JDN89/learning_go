package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/JDN89/learning_go/internal/assets"
	"github.com/JDN89/learning_go/internal/server"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

func main() {
	// Declare an instance of the config struct.
	var cfg config
	// Read the value of the port and env command-line flags into the config struct. We
	// default to using the port number 4000 and the environment "development" if no
	// corresponding flags are provided.
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()
	// Initialize a new structured logger which writes log entries to the standard out
	// stream.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// Declare an instance of the application struct, containing the config struct and
	// the logger.

	// Initialize a new htmlRenderer, parsing the base template and all partial
	// templates from assets/html into the shared template set.
	htmlRenderer, err := server.NewHTMLRenderer(assets.HTMLFiles, "base.tmpl", "partials/*.tmpl")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &Application{
		config: cfg,
		logger: logger,
		html:   htmlRenderer,
	}
	// Declare a new servemux and add a /v1/healthcheck route which dispatches requests
	// to the healthcheckHandler method (which we will create in a moment).
	mux := http.NewServeMux()
	// Serve embedded CSS, JavaScript, and other static assets from /static/.
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(assets.StaticFiles))))
	mux.Handle("/healthz", Healthcheck(app.config.env, version))
	mux.HandleFunc("/{$}", app.Home)

	// mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	// Declare an HTTP server which listens on the port provided in the config struct,
	// uses the servemux we created above as the handler, has some sensible timeout
	// settings, and writes any log messages to the structured logger at Error level.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}
	// Start the HTTP server.
	logger.Info("starting server", "addr", srv.Addr, "env", app.config.env)
	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
