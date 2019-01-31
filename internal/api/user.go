package api

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	ID    uint64 `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Error string `json:"error,omitempty"`
}
