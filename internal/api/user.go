package api

import (
	"encoding/json"
	"net/http"

	"github.com/tinrab/kit/id"

	"github.com/tinrab/event-source-scream/internal/pkg/command"
	userCommand "github.com/tinrab/event-source-scream/internal/user/command"
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

	data := userCommand.CreateUser{
		Name: req.Name,
	}
	cmd := command.New(s.idGenerator.Generate(), userCommand.KindCreateUser, data)
	var res command.Result

	if err := s.bus.PublishCommand(cmd, &res, timeout); err != nil {
		s.respond(w, http.StatusBadRequest, CreateUserResponse{
			Error: err.Error(),
		})
		return
	}

	if res.IsError() {
		s.respond(w, http.StatusBadRequest, CreateUserResponse{
			Error: res.Error,
		})
		return
	}

	var resultData userCommand.CreateUserResult
	if err := json.Unmarshal(res.Data, &resultData); err != nil {
		s.respond(w, http.StatusBadRequest, CreateUserResponse{
			Error: res.Error,
		})
	}

	s.respond(w, http.StatusOK, CreateUserResponse{
		ID:   resultData.ID,
		Name: resultData.Name,
	})
}
