package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/etiennemarais/go-openapi-swag-demo/pkg/router"
)

// Server holds our server dependencies
type Server struct {
	mux http.Handler
}

func main() {
	chain, err := router.New(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := Server{chain}
	s.Start()
}

// Start will start the http server
func (s Server) Start() error {
	return http.ListenAndServe(":3000", s.mux)
}
