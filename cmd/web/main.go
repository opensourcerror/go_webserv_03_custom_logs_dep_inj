package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// go run	 ./cmd/web -addr=":40000"

type application struct {
	logger *slog.Logger
}

func main() {
	// :4000 will be default value if no flag
	addr := flag.String("addr", ":4000", "HTTP net addr")
	flag.Parse() //this must be before using `addr`

	// nil when u want to use the default settings
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	mux.HandleFunc("/sb", app.sb)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// addr is actually a pointer, and we need to dereference it
	//( prefix it withthe * symbol) before using it
	// args ...any
	logger.Info("srv up on ", "addr", *addr)
	// log.Printf("srv up on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
	// log.Fatal(err)
}
