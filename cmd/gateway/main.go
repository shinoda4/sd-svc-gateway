package main

import (
	"log"

	"github.com/shinoda4/sd-svc-gateway/internal/config"
	"github.com/shinoda4/sd-svc-gateway/internal/router"
	grpc_transport "github.com/shinoda4/sd-svc-gateway/internal/transport/grpc"
	http_transport "github.com/shinoda4/sd-svc-gateway/internal/transport/http"
)

func main() {
	cfg := config.Load()

	// Initialize Auth gRPC client
	authClient, authConn := grpc_transport.NewAuthClient(cfg.AuthSvcURL)
	defer authConn.Close()

	// Initialize Auth HTTP handler
	authH := http_transport.NewAuthHandler(authClient)

	r := router.Setup(cfg, authH)

	log.Printf("Gateway starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start gateway: %v", err)
	}
}
