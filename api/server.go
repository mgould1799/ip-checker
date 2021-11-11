package api

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/mgould1799/ip-checker/ipChecker"
)

type Server struct {
	port int 
	ipCheck *ipChecker.IpChecker
}

func NewServer(dbPath string, port int) *Server{
	return &Server{port: port, ipCheck: ipChecker.NewIpChecker(dbPath)}
}

func (s Server) Start() {
	var wait time.Duration
    flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
    flag.Parse()

    r := mux.NewRouter()
    // Add your routes as needed

	r.HandleFunc("/check-ip", s.checkIpInWhiteList).Methods("GET")
	r.HandleFunc("/healthz", healthz).Methods("GET")
	http.Handle("/", r)

    srv := &http.Server{
        Addr:         fmt.Sprintf("0.0.0.0:%d", s.port),
        // Good practice to set timeouts to avoid Slowloris attacks.
        WriteTimeout: time.Second * 15,
        ReadTimeout:  time.Second * 15,
        IdleTimeout:  time.Second * 60,
        Handler: r, // Pass our instance of gorilla/mux in.
    }

    // Run our server in a goroutine so that it doesn't block.
    go func() {
		log.Infof("starting server on port %d", s.port)
        if err := srv.ListenAndServe(); err != nil {
            log.Println(err)
        }
    }()

    c := make(chan os.Signal, 1)
    // We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
    // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
    signal.Notify(c, os.Interrupt)

    // Block until we receive our signal.
    <-c

    // Create a deadline to wait for.
    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer cancel()
    // Doesn't block if no connections, but will otherwise wait
    // until the timeout deadline.
    srv.Shutdown(ctx)
    // Optionally, you could run srv.Shutdown in a goroutine and block on
    // <-ctx.Done() if your application should wait for other services
    // to finalize based on context cancellation.
	s.Stop()
    os.Exit(0)
}

func (s Server) Stop() {
	s.ipCheck.StopIpChecker()
	log.Info("shut down IP checker")
	log.Info("shutting server down")
}

