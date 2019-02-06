package event

const (
	KindUserCreated = "user.created"
)

type UserCreated struct {
	Name string `json:"name"`
}
