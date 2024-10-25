// Code generated by goa v3.19.1, DO NOT EDIT.
//
// poc endpoints
//
// Command:
// $ goa gen goa_poc/goa/design

package poc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "poc" service endpoints.
type Endpoints struct {
	Get goa.Endpoint
}

// NewEndpoints wraps the methods of the "poc" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Get: NewGetEndpoint(s),
	}
}

// Use applies the given middleware to all the "poc" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Get = m(e.Get)
}

// NewGetEndpoint returns an endpoint function that calls the method "get" of
// service "poc".
func NewGetEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetPayload)
		return s.Get(ctx, p)
	}
}
