// account-management/cmd/server/main.go
package main

import (
	"account-management/internal/config"
	"account-management/internal/http"
	"account-management/internal/repo"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.Load()
	db, err := repo.NewMgmtDB(cfg.MgmtDBUrl)
	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}
	defer db.Close()

	r := gin.Default()
	http.RegisterRoutes(r, db, cfg)

	log.Printf("mgmt service listening on %s", cfg.ListenAddr)
	if err := r.Run(cfg.ListenAddr); err != nil {
		log.Fatal(err)
	}
}
