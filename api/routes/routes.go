package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/kanevidzro/gokernel/api/handlers"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	// Healthcheck
	r.GET("/health", handlers.HealthCheck)

	// User handler
	userHandler := handlers.NewUserHandler(db)

	v1 := r.Group("/v1")
	{
		v1.GET("/users/:id", userHandler.GetUser)
	}
}
