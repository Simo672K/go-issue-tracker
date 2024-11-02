package routes

import (
	"context"
	"net/http"

	"github.com/Simo672K/issue-tracker/pkg/router"
	"github.com/Simo672K/issue-tracker/service"
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

	//* Mounting User routes
	UserRoutes(router)

	router.GET("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	router.GET("/api/v1/email-test", func(w http.ResponseWriter, r *http.Request) {
		emailContent := &service.EmailContent{
			Subject: "Wellcome email",
			Content: "Hello there, this is an automatic email for wellcoming you!",
		}
		if err := service.EmailService([]string{"<reciever-email>"}, emailContent); err != nil {
			http.Error(w, "Failed to send email : "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Email sent successfully!"))
	})
	return nil
}
