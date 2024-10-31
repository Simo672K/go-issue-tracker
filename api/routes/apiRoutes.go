package routes

import (
	"context"
	"net/http"

	"github.com/Simo672K/issue-tracker/pkg/router"
)

/**
* 	Api structure
*   /api/<version>/<resource>/<create|delete|update|all|id>
 */
func MuxRouter(ctx context.Context, mux *http.ServeMux) error {
	router := router.NewRouter(ctx, mux)

	//* Mounting Authentication routes
	AuthRoutes(router)

	//* Mounting Project routes
	ProjectRoutes(router)

	router.GET("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return nil
}
