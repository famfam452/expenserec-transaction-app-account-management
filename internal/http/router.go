// account-management/internal/http/router.go
package http

import (
	"account-management/internal/config"
	"account-management/internal/handlers"
	"account-management/internal/middleware"
	"account-management/internal/repo"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *repo.MgmtDB, cfg config.Config) {
	v1 := r.Group("/api-v1/account-management")
	authMW := middleware.JWTAuth(cfg.JWTSecret)

	h := handlers.NewAccountsHandler(db)

	v1.Use(authMW)
	v1.GET("/list-account", h.List)
	v1.PUT("/create-account", h.Create)
	v1.GET("/account/:account_id", h.Get)
}
