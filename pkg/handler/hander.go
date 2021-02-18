package handler

import (
	"context"
)

// Handler wraps our specific handlers.
type Handler struct {
	Demo DemoHandler
}

// DemoHandler contains our user specific handlers.
type DemoHandler struct{}

func New(ctx context.Context) (Handler, error) {
	return Handler{
		Demo: DemoHandler{},
	}, nil
}
