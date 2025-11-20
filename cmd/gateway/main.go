/*
 * Copyright (c) 2025-11-20 shinoda4
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"log"

	authpb "github.com/shinoda4/sd-grpc-proto/proto/auth/v1"
	"github.com/shinoda4/sd-svc-gateway/internal/config"
	"github.com/shinoda4/sd-svc-gateway/internal/router"
	grpctransport "github.com/shinoda4/sd-svc-gateway/internal/transport/grpc"
	httptransport "github.com/shinoda4/sd-svc-gateway/internal/transport/http/auth"
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
	authH := httptransport.NewHandler(authClient)

	r := router.Setup(cfg, authH)

	log.Printf("Gateway starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start gateway: %v", err)
	}
}
