package routes

import (
	"github.com/Simo672K/issue-tracker/api/handler"
	"github.com/Simo672K/issue-tracker/api/middleware"
	"github.com/Simo672K/issue-tracker/pkg/router"
)

func ProjectRoutes(pr *router.Router) {
	pr.GET("/api/v1/project/list", handler.ListAssociatedProjectsHandler, middleware.AuthMiddleware)
	pr.POST("/api/v1/project/new", handler.CreateProjectHandler, middleware.AuthMiddleware)
	pr.GET("/api/v1/project/{projectId}", handler.ProjectInfoHandler, middleware.AuthMiddleware)
	pr.GET("/api/v1/project/{projectId}/status", handler.AuthSignUpHandler, middleware.AuthMiddleware)
	pr.GET("/api/v1/project/{projectId}/delete", handler.AuthSignUpHandler, middleware.AuthMiddleware)

	// TODO: Issues associated with a project
	pr.GET("/api/v1/project/{projectId}/issue", handler.AuthSignUpHandler, middleware.AuthMiddleware)
	pr.POST("/api/v1/project/{projectId}/issue/new", handler.AuthSignUpHandler, middleware.AuthMiddleware)
	pr.GET("/api/v1/project/{projectId}/issue/{issueId}", handler.AuthSignUpHandler, middleware.AuthMiddleware)
	pr.POST("/api/v1/project/{projectId}/issue/{issueId}/assign", handler.AuthSignUpHandler, middleware.AuthMiddleware)
	pr.GET("/api/v1/project/{projectId}/issue/{issueId}/status", handler.AuthSignUpHandler, middleware.AuthMiddleware)

	// TODO: Managers associated with a project
	pr.GET("/api/v1/project/{projectId}/manager", handler.AuthSignUpHandler, middleware.AuthMiddleware)
	pr.POST("/api/v1/project/{projectId}/manager/add", handler.AuthSignUpHandler, middleware.AuthMiddleware)
	pr.GET("/api/v1/project/{projectId}/manager/{managerId}/delete", handler.AuthSignUpHandler, middleware.AuthMiddleware)

	// TODO: Developers associated with a project
	pr.GET("/api/v1/project/{projectId}/developer", handler.AuthSignUpHandler, middleware.AuthMiddleware)
	pr.POST("/api/v1/project/{projectId}/developer/add", handler.AuthSignUpHandler, middleware.AuthMiddleware)
	pr.GET("/api/v1/project/{projectId}/developer/{developerId}/delete", handler.AuthSignUpHandler, middleware.AuthMiddleware)
}
