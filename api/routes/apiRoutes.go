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
	// auth routes
	mux.HandleFunc("POST /api/v1/auth/signup", handler.AuthSignUpHandler)
	mux.HandleFunc("POST /api/v1/auth/signin", handler.AuthSignInHandler)
	// mux.HandleFunc("POST /api/v1/auth/change-password", handler.AuthSignUpHandler)
	// mux.HandleFunc("POST /api/v1/auth/reset-password", handler.AuthSignUpHandler)

	// mux.HandleFunc("GET /api/v1/verify-email", handler.AuthSignUpHandler)

	return nil
}
