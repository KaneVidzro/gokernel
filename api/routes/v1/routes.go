package v1

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kanevidzro/gokernel/internal/admin"
	"github.com/kanevidzro/gokernel/internal/auth"
	"github.com/kanevidzro/gokernel/internal/user"
	"github.com/kanevidzro/gokernel/pkg/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
)

func Register(r *gin.RouterGroup, db *sql.DB, redis *redis.Client, cfg *config.Config) {
    // Health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    // Prometheus metrics
    r.GET("/metrics", gin.WrapH(promhttp.Handler()))

    // Repositories and services
    userRepo := &user.Repository{DB: db}
    authService := &auth.AuthService{
        UserRepo:  userRepo,
        Redis:     redis,
        JWTSecret: []byte(cfg.JWTSecret),
    }
    authHandler := &auth.AuthHandler{Service: authService}
    userHandler := &user.HandlerV1{Repository: userRepo}

    // Auth routes
    authGroup := r.Group("/auth")
    authGroup.Use(auth.RateLimitMiddleware(redis, 15, time.Minute)) // 15 requests per minute
    {
        authGroup.POST("/register", authHandler.Register)
        authGroup.POST("/login", authHandler.Login)
        authGroup.POST("/logout", authHandler.Logout)
    }

    // Protected routes
    authMiddleware := auth.AuthRequired([]byte(cfg.JWTSecret))

    protected := r.Group("/")
    protected.Use(authMiddleware)
    {
        protected.GET("/users/:id", userHandler.GetUser)
    }


adminService := &admin.Service{UserRepo: userRepo}
adminHandler := &admin.Handler{Service: adminService}

admin := r.Group("/admin")
admin.Use(auth.AuthRequired([]byte(cfg.JWTSecret)), auth.RequireRole("admin"))
{
    admin.GET("/dashboard", adminHandler.Dashboard)
    admin.GET("/users", adminHandler.ListUsers)
    admin.PUT("/users/:id/role", adminHandler.UpdateRole)
    admin.POST("/users/:id/deactivate", adminHandler.DeactivateUser)
}


}
