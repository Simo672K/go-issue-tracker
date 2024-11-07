package middleware

import (
	"context"
	"net/http"

	"github.com/Simo672K/issue-tracker/internal/auth"
	"github.com/Simo672K/issue-tracker/pkg/router"
	"github.com/Simo672K/issue-tracker/utils"
)

func WithRoleMiddlware(role auth.Role) router.Middleware {
	return func(ctx *context.Context, handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(*ctx)
			profileId := r.Context().Value("profileId")
			projectId := r.PathValue("projectId")

			per := auth.NewPermission(role)
			if per.HasAccessTo(*ctx, profileId.(string), projectId) {
				handler.ServeHTTP(w, r)
				return
			}

			utils.WriteJsonError(
				w,
				http.StatusUnauthorized,
				"UNAUTHORIZED",
				"Access denied",
			)
		})
	}
}
