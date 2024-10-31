package routes

import (
	"github.com/Simo672K/issue-tracker/api/handler"
	"github.com/Simo672K/issue-tracker/pkg/router"
)

func ProjectRoutes(pr *router.Router) {

	pr.GET("/api/v1/project/list", handler.AuthSignUpHandler)
	pr.POST("/api/v1/project/new", handler.AuthSignUpHandler)
	pr.GET("/api/v1/project/{projectId}", handler.AuthSignUpHandler)
	pr.GET("/api/v1/project/{projectId}/status", handler.AuthSignUpHandler)
	pr.GET("/api/v1/project/{projectId}/delete", handler.AuthSignUpHandler)

	//! Issues associated with a project
	pr.GET("/api/v1/project/{projectId}/issue", handler.AuthSignUpHandler)
	pr.POST("/api/v1/project/{projectId}/issue/new", handler.AuthSignUpHandler)
	pr.GET("/api/v1/project/{projectId}/issue/{issueId}", handler.AuthSignUpHandler)
	pr.POST("/api/v1/project/{projectId}/issue/{issueId}/assign", handler.AuthSignUpHandler)
	pr.GET("/api/v1/project/{projectId}/issue/{issueId}/status", handler.AuthSignUpHandler)

	//! Managers associated with a project
	pr.GET("/api/v1/project/{projectId}/manager", handler.AuthSignUpHandler)
	pr.POST("/api/v1/project/{projectId}/manager/add", handler.AuthSignUpHandler)
	pr.GET("/api/v1/project/{projectId}/manager/{managerId}/delete", handler.AuthSignUpHandler)

	//! Developers associated with a project
	pr.GET("/api/v1/project/{projectId}/developer", handler.AuthSignUpHandler)
	pr.POST("/api/v1/project/{projectId}/developer/add", handler.AuthSignUpHandler)
	pr.GET("/api/v1/project/{projectId}/developer/{developerId}/delete", handler.AuthSignUpHandler)
}
