package nrmux

import (
	"context"

	"github.com/etiennemarais/go-openapi-swag-demo/pkg/nrhttptreemux"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func nilmux() *nrhttptreemux.Router {
	app, _ := newrelic.NewApplication(
		newrelic.ConfigEnabled(false),
	)

	return nrhttptreemux.New(app)
}

// New will create a new httptreemux.ContextMux, automatically configuring
// instrumentation with New Relic. It's important to note that if we cannot
// create the New Relic application a warning or error will be logged but a
// valid mux will still be returned. This allows us to work locally without
// having setup any credentials.
func New(ctx context.Context) *nrhttptreemux.Router {
	// Demo mux
	return nilmux()
}
