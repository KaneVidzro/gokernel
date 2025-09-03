package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	v1 "github.com/kanevidzro/gokernel/api/routes/v1"
	v2 "github.com/kanevidzro/gokernel/api/routes/v2"
	"github.com/kanevidzro/gokernel/pkg/config"
	"github.com/redis/go-redis/v9"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB, redis *redis.Client, cfg *config.Config) {
    v1Group := r.Group("/v1")
    v1.Register(v1Group, db, redis, cfg)

    v2Group := r.Group("/v2")
    v2.Register(v2Group)
}
