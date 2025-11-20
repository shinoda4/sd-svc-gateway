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

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shinoda4/sd-svc-gateway/internal/config"
	"github.com/shinoda4/sd-svc-gateway/internal/middleware"
	http_transport "github.com/shinoda4/sd-svc-gateway/internal/transport/http/auth"
)

func Setup(cfg *config.Config, authH *http_transport.Handler) *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Public routes (Auth service)
	auth := r.Group("/auth")
	{
		auth.POST("/register", authH.Register)
		auth.POST("/login", authH.Login)
		auth.GET("/verify", authH.VerifyEmail)
		auth.POST("/forgot-password", authH.ForgotPassword)
		auth.POST("/reset-password", authH.ResetPassword)
		auth.POST("/refresh", authH.RefreshToken)
		auth.POST("/logout", authH.Logout)
	}

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(cfg))
	{
		// Example protected route
		api.GET("/me", func(c *gin.Context) {
			// We can also use the auth handler's Me method if we want to proxy it
			// But usually /me is handled by the auth service too.
			// Let's route it to auth handler as well, but maybe via /auth/me or /api/me?
			// The proto has Me rpc.
			authH.Me(c)
		})

		// Add more protected service proxies here
	}

	return r
}
