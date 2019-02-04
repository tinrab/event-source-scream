package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/tinrab/event-source-scream/internal/pkg/bus"
)

type Server struct {
	port   uint16
	bus    *bus.Bus
	router *mux.Router
}

const (
	timeout = 5 * time.Second
)

func NewServer(port uint16, b *bus.Bus) *Server {
	return &Server{
		port:   port,
		bus:    b,
		router: mux.NewRouter(),
	}
}

func (s *Server) Run() error {
	s.router.HandleFunc("/users", s.createUser).
		Methods("POST")
	s.router.HandleFunc("/users/{userId}/screams", s.listScreamsByUser).
		Methods("GET")
	s.router.HandleFunc("/screams", s.createScream).
		Methods("POST")

	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router)
}

func (s *Server) respond(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Print(err)
	}
}
