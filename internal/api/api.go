package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/tinrab/event-source-store/internal/pkg/bus"
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

	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router)
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.respond(w, http.StatusBadRequest, CreateUserResponse{
			Error: err.Error(),
		})
		return
	}

	var res CreateUserResponse

	if err := s.bus.Request("api.user.create", req, &res, timeout); err != nil {
		s.respond(w, http.StatusBadRequest, CreateUserResponse{
			Error: err.Error(),
		})
		return
	}

	if len(res.Error) != 0 {
		s.respond(w, http.StatusBadRequest, res)
		return
	}

	s.respond(w, http.StatusOK, res)
}

func (s *Server) respond(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Print(err)
	}
}
