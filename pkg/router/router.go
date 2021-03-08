package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/etiennemarais/go-openapi-swag-demo/pkg/handler"
	"github.com/etiennemarais/go-openapi-swag-demo/pkg/nrmux"
	"github.com/justinas/alice"
)

// New will create a new http handler containing our routes.
//
// @title Demo API
// @description The demo API handles all calls to the demo namespace for purposes of illustration.
//
// @contact.name Etienne Marais
// @contact.url https://<company>.zendesk.com/hc/en-us
//
// @host api.swag.com
// @BasePath /
// @query.collection.format multi
// @schemes https
//
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name API-Token
func New(ctx context.Context) (http.Handler, error) {
	h, err := handler.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("create handlers: %w", err)
	}

	router := nrmux.New(ctx)

	router.GET("/demo", h.Demo.ListDemo)
	router.POST("/demo/:id", h.Demo.CreateDemo)

	chain := alice.New()

	return chain.Then(router), nil
}
