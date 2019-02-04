package api

import (
	"encoding/json"
	"net/http"

	"github.com/tinrab/kit/id"
)

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	ID    id.ID  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Error string `json:"error,omitempty"`
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

	if err := s.bus.Request("user.create", req, &res, timeout); err != nil {
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
