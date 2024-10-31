package routes

import (
	"github.com/Simo672K/issue-tracker/api/handler"
	"github.com/Simo672K/issue-tracker/api/middleware"
	"github.com/Simo672K/issue-tracker/pkg/router"
)

func AuthRoutes(ar *router.Router) {

	ar.POST("/api/v1/auth/signup", handler.AuthSignUpHandler)
	ar.POST("/api/v1/auth/signin", handler.AuthSignInHandler)
	ar.POST("/api/v1/auth/change-password", handler.AuthSignUpHandler)
	ar.POST("/api/v1/auth/reset-password", handler.AuthSignUpHandler)

	ar.GET("/api/v1/test-auth", handler.AuthTest, middleware.AuthMiddleware)
	ar.GET("/api/v1/verify-email", handler.AuthSignUpHandler)

}
