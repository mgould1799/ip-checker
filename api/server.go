package api

import (
	"fmt"
	"net/http"

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
	r := mux.NewRouter()
	r.HandleFunc("/check-ip", s.checkIp).Methods("GET")
	r.HandleFunc("/healthz", healthz).Methods("GET")
	http.Handle("/", r)


	http.ListenAndServe(fmt.Sprintf(":%d", s.port), r)
}

