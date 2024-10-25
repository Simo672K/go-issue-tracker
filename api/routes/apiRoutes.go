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
	//! Auth routes
	mux.HandleFunc("POST /api/v1/auth/signup", handler.AuthSignUpHandler)
	mux.HandleFunc("POST /api/v1/auth/signin", handler.AuthSignInHandler)
	mux.HandleFunc("GET /api/v1/test-auth", handler.AuthTest)

	// mux.HandleFunc("POST /api/v1/auth/change-password", handler.AuthSignUpHandler)
	// mux.HandleFunc("POST /api/v1/auth/reset-password", handler.AuthSignUpHandler)

	// mux.HandleFunc("GET /api/v1/verify-email", handler.AuthSignUpHandler)

	//! Project handling routes
	mux.HandleFunc("GET /api/v1/project/list", handler.AuthSignUpHandler)
	mux.HandleFunc("POST /api/v1/project/create", handler.AuthSignUpHandler)
	mux.HandleFunc("GET /api/v1/project/{projectId}", handler.AuthSignUpHandler)
	mux.HandleFunc("GET /api/v1/project/{projectId}/status", handler.AuthSignUpHandler)
	mux.HandleFunc("GET /api/v1/project/{projectId}/delete", handler.AuthSignUpHandler)

	//! Issues associated with a project
	mux.HandleFunc("GET /api/v1/project/{projectId}/issue", handler.AuthSignUpHandler)
	mux.HandleFunc("GET /api/v1/project/{projectId}/issue/add", handler.AuthSignUpHandler)
	mux.HandleFunc("GET /api/v1/project/{projectId}/issue/{issueId}", handler.AuthSignUpHandler)
	mux.HandleFunc("GET /api/v1/project/{projectId}/issue/{issueId}/assign", handler.AuthSignUpHandler)
	mux.HandleFunc("GET /api/v1/project/{projectId}/issue/{issueId}/status", handler.AuthSignUpHandler)

	//! Managers associated with a project
	mux.HandleFunc("GET /api/v1/project/{projectId}/manager", handler.AuthSignUpHandler)
	mux.HandleFunc("GET /api/v1/project/{projectId}/manager/add", handler.AuthSignUpHandler)
	mux.HandleFunc("GET /api/v1/project/{projectId}/manager/{managerId}/delete", handler.AuthSignUpHandler)

	//! Developers associated with a project
	mux.HandleFunc("GET /api/v1/project/{projectId}/developer", handler.AuthSignUpHandler)
	mux.HandleFunc("GET /api/v1/project/{projectId}/developer/add", handler.AuthSignUpHandler)
	mux.HandleFunc("GET /api/v1/project/{projectId}/developer/{developerId}/delete", handler.AuthSignUpHandler)

	return nil
}
