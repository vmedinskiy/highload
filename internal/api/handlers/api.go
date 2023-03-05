package handlers

import (
	"sync"

	api "github.com/vmedinskiy/highload/api/generated"
)

// Compile-time check for Handler.
var _ api.Handler = (*ServerHandler)(nil)

type ServerHandler struct {
	users map[string]api.User
	id    string
	mux   sync.Mutex
}

func NewServerHandler() *ServerHandler {
	return &ServerHandler{
		users: make(map[string]api.User),
	}
}
