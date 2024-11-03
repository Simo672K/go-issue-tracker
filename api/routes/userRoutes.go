package routes

import (
	"github.com/Simo672K/issue-tracker/api/handler"
	"github.com/Simo672K/issue-tracker/pkg/router"
)

func UserRoutes(ur *router.Router) {
	ur.GET("/api/v1/verify-email/{userId}", handler.VerifyUserEmail)
}
