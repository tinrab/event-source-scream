package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tinrab/kit/id"
)

type CreateScreamRequest struct {
	UserID id.ID  `json:"user_id"`
	Body   string `json:"body"`
}

type CreateScreamResponse struct {
	ID        id.ID      `json:"id,omitempty"`
	UserID    id.ID      `json:"user_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Body      string     `json:"body,omitempty"`
	Error     string     `json:"error,omitempty"`
}

type ListScreamsRequest struct {
	UserID id.ID `json:"user_id"`
}

type ListScreamsResponse struct {
	UserID  id.ID            `json:"user_id,omitempty"`
	Screams []ScreamResponse `json:"screams,omitempty"`
	Error   string           `json:"error,omitempty"`
}

type ScreamResponse struct {
	ID        id.ID      `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Body      string     `json:"body,omitempty"`
}

func (s *Server) createScream(w http.ResponseWriter, r *http.Request) {
	var req CreateScreamRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.respond(w, http.StatusBadRequest, CreateScreamResponse{
			Error: err.Error(),
		})
		return
	}

	if userID, err := id.ParseID(r.Header.Get("X-USER-ID")); err != nil {
		s.respond(w, http.StatusBadRequest, CreateScreamResponse{
			Error: "invalid user ID",
		})
		return
	} else {
		req.UserID = userID
	}

	var res CreateScreamResponse

	if err := s.bus.Request("scream.create", req, &res, timeout); err != nil {
		s.respond(w, http.StatusBadRequest, CreateScreamResponse{
			Error: err.Error(),
		})
		return
	}

	if len(res.Error) != 0 {
		s.respond(w, http.StatusBadRequest, CreateScreamResponse{
			Error: res.Error,
		})
		return
	}

	s.respond(w, http.StatusOK, res)
}

func (s *Server) listScreamsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := id.ParseID(vars["userId"])
	if err != nil {
		s.respond(w, http.StatusBadRequest, ListScreamsResponse{
			Error: err.Error(),
		})
		return
	}

	req := ListScreamsRequest{
		UserID: userID,
	}
	var res ListScreamsResponse

	if err := s.bus.Request("scream.list", req, &res, timeout); err != nil {
		s.respond(w, http.StatusBadRequest, ListScreamsResponse{
			Error: err.Error(),
		})
	}

	if len(res.Error) != 0 {
		s.respond(w, http.StatusBadRequest, ListScreamsResponse{
			Error: res.Error,
		})
		return
	}

	s.respond(w, http.StatusOK, res)
}
