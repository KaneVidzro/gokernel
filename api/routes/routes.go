package routes

import (
	"database/sql"

	"github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
	"github.com/kanevidzro/gokernel/api/handlers"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB, redis *redis.Client)  {

	    h := &handlers.Handler{
        DB:    db,
        Redis: redis,
    }
	// Healthcheck
	r.GET("/health", h.HealthCheck)

	// User handler
	userHandler := handlers.NewUserHandler(db)

	v1 := r.Group("/v1")
	{
		v1.GET("/users/:id", userHandler.GetUser)
	}
}
