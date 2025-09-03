// api/handlers/health_handler.go
package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Handler struct {
    DB    *sql.DB
    Redis *redis.Client
}

func (h *Handler) HealthCheck(c *gin.Context) {
    dbErr := h.DB.Ping()
    redisErr := h.Redis.Ping(context.Background()).Err()

    status := gin.H{"status": "ok"}

    if dbErr != nil {
        status["db"] = "down"
    }
    if redisErr != nil {
        status["redis"] = "down"
    }

    c.JSON(http.StatusOK, status)
}
