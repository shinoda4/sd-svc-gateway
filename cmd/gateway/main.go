package main

import (
	"context"
	"log"

	authpb "github.com/shinoda4/sd-grpc-proto/proto/auth/v1"
	"github.com/shinoda4/sd-svc-gateway/internal/config"
	"github.com/shinoda4/sd-svc-gateway/internal/router"
	grpctransport "github.com/shinoda4/sd-svc-gateway/internal/transport/grpc"
	httptransport "github.com/shinoda4/sd-svc-gateway/internal/transport/http"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	// Initialize Auth gRPC client
	authClient, authConn := grpctransport.NewAuthClient(cfg.AuthSvcURL)
	defer func(authConn *grpc.ClientConn) {
		err := authConn.Close()
		if err != nil {
			return
		}
	}(authConn)

	healthCheckResponse, err := authClient.HealthCheck(context.Background(), &authpb.HealthCheckRequest{
		Message: "hello",
	})
	if err != nil {
		return
	}

	log.Printf("health check response: %v", healthCheckResponse)

	// Initialize Auth HTTP handler
	authH := httptransport.NewAuthHandler(authClient)

	r := router.Setup(cfg, authH)

	log.Printf("Gateway starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start gateway: %v", err)
	}
}
