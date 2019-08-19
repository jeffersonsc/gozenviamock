package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jeffersonsc/gozenviamock/pkg/mutiple"
	"github.com/jeffersonsc/gozenviamock/pkg/single"
	"github.com/urfave/negroni"

	"github.com/google/subcommands"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server .
type Server struct {
	port string
}

// Name .
func (*Server) Name() string { return "server" }

// Synopsis .
func (*Server) Synopsis() string { return "Starter new server" }

// Usage .
func (*Server) Usage() string {
	return `server [-port] <some text>:
  Stater new server
`
}

// SetFlags .
func (s *Server) SetFlags(f *flag.FlagSet) {
	f.StringVar(&s.port, "port", "3000", "port usage system")
}

// Execute .
func (s *Server) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	n := negroni.Classic()
	n.Use(applicationJSON())
	n.Use(basicAuth())

	r := mux.NewRouter()
	n.UseHandler(r)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("IT'S WORK!"))
	}).Methods(http.MethodGet)

	single.RegisterRouter(r)
	mutiple.RegisterRouter(r)

	cerr := make(chan error, 1)

	// Gracefull shutting down
	go func(ctx context.Context) {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint
		log.Println("HTTP server shutting down")
		cerr <- nil
	}(ctx)

	// Starter server
	go func(ctx context.Context) {
		log.Printf("HTTP server listening on %v", s.port)
		// Enable cors
		cors := handlers.AllowedOrigins([]string{"*"})
		if err := http.ListenAndServe(fmt.Sprintf(":%s", s.port), handlers.CORS(cors)(n)); err != nil {
			cerr <- err
		}
	}(ctx)

	// Waiting
	err := <-cerr
	if err != nil {
		log.Println(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func applicationJSON() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	})
}

func basicAuth() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if r.URL.Path == "/" {
			next(w, r)
			return
		}

		user, pass, ok := r.BasicAuth()
		if !ok || user != "user" || pass != "pass" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, `{"error":"unauthorized"}`)
			return
		}
		next(w, r)
	})
}
