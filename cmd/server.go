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
	r := mux.NewRouter()
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
		if err := http.ListenAndServe(fmt.Sprintf(":%s", s.port), handlers.CORS(cors)(r)); err != nil {
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
