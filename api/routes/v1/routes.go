package v1

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/kanevidzro/gokernel/internal/auth"
	"github.com/kanevidzro/gokernel/internal/user"
	"github.com/kanevidzro/gokernel/pkg/config"
	"github.com/redis/go-redis/v9"
)

func Register(r *gin.RouterGroup, db *sql.DB, redis *redis.Client, cfg *config.Config) {
    // Health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

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
}
