package oapi

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	oapiMiddleware "github.com/oapi-codegen/nethttp-middleware"
)

func NewChiRouter(ssi StrictServerInterface) (*chi.Mux, error) {
	swagger, err := GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("error loading swagger spec:\n%s", err)
	}

	// Setup our Chi Router
	r := chi.NewRouter()
	r.Use(oapiMiddleware.OapiRequestValidator(swagger))
	r.Use(middleware.Logger)

	// Register our
	strictHandler := NewStrictHandler(ssi, nil)
	HandlerFromMux(strictHandler, r)

	return r, nil
}
