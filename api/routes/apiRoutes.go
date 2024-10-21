package routes

import (
	"context"
	"net/http"

	"github.com/Simo672K/issue-tracker/api/handler"
)

/**
* 	Api structure
*   /api/<version>/<resource>/<create|delete|update|all|id>
 */
func MuxRouter(ctx context.Context, mux *http.ServeMux) error {
	mux.HandleFunc("GET /checkhealth", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"message\":\"Hello world\"}"))
	})
	mux.HandleFunc("POST /api/v1/users/create", handler.CreateUserHandler)

	return nil
}
