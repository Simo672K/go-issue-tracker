package router

import (
	"context"
	"net/http"
)

type Middleware func(context.Context, http.Handler) http.Handler
type Controller func(http.ResponseWriter, *http.Request)

type Router struct {
	Ctx context.Context
	Mux *http.ServeMux
}

// TODO: making router function that initialize a router based on a given mux
func NewRouter(ctx context.Context, mux *http.ServeMux) *Router {
	return &Router{
		Ctx: ctx,
		Mux: mux,
	}
}

// TODO: create a function that accepts a route and handlers, and pipe thous handlers into each other internaly,
// TODO: commonly khowns as middlwares and controllers
// Handle method for flexible HTTP method handling with middleware chaining
func (r *Router) Handle(method, path string, controller Controller, middlewares ...Middleware) {
	// Base handler for the controller
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Check if the request method matches
		if req.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		// Call the controller if the method matches
		controller(w, req)
	})

	// Apply middlewares to the handler
	pipedHandler := applyMiddlewaresPipe(r.Ctx, handler, middlewares...)
	r.Mux.Handle(path, pipedHandler)
}

// Specific HTTP method helpers (GET, POST, etc.) for convenience
func (r *Router) GET(path string, controller Controller, middlewares ...Middleware) {
	r.Handle("GET", path, controller, middlewares...)
}

func (r *Router) POST(path string, controller Controller, middlewares ...Middleware) {
	r.Handle("POST", path, controller, middlewares...)
}

func (r *Router) PUT(path string, controller Controller, middlewares ...Middleware) {
	r.Handle("PUT", path, controller, middlewares...)
}

func (r *Router) DELETE(path string, controller Controller, middlewares ...Middleware) {
	r.Handle("DELETE", path, controller, middlewares...)
}

// * this function applys middlewares recursively, and the middlewares order maters
func applyMiddlewaresPipe(ctx context.Context, handler http.Handler, middlewares ...Middleware) http.Handler {
	if len(middlewares) == 1 {
		return middlewares[0](ctx, handler)
	}
	// extracts the current middlware
	currMiddlwr := middlewares[len(middlewares)-1]
	// separates other middlware for piping
	mdlwrs := middlewares[:len(middlewares)-1]
	return currMiddlwr(ctx, applyMiddlewaresPipe(ctx, handler, mdlwrs...))
}