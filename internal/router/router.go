package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shinoda4/sd-svc-gateway/internal/config"
	"github.com/shinoda4/sd-svc-gateway/internal/middleware"
	"github.com/shinoda4/sd-svc-gateway/internal/proxy"
)

func Setup(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Public routes (Auth service)
	// Forward everything starting with /auth to the auth service
	// The auth service likely expects /api/v1/... or similar, but let's assume it handles its own routing.
	// If we want to strip the prefix, we might need adjustment, but usually a gateway forwards the path.
	// Let's assume we forward /auth/*path to AuthSvcURL/*path
	// Or if Auth service is at root, we might need to strip /auth.
	// Let's assume for now we just forward.

	authProxy := proxy.NewReverseProxy(cfg.AuthSvcURL)
	r.Any("/auth/*proxyPath", authProxy)

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(cfg))
	{
		// Example protected route
		api.GET("/me", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			email, _ := c.Get("email")
			c.JSON(200, gin.H{
				"message": "You are authenticated",
				"userID":  userID,
				"email":   email,
			})
		})

		// Add more protected service proxies here
	}

	return r
}
