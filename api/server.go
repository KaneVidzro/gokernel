package api

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/kanevidzro/gokernel/api/routes"
	"github.com/kanevidzro/gokernel/pkg/config"
	"github.com/kanevidzro/gokernel/pkg/middleware"
)

type Server struct {
	engine *gin.Engine
	logger *zap.Logger
	cfg    *config.Config
	db     *sql.DB
	redis  *redis.Client
}

func NewServer(logger *zap.Logger, cfg *config.Config) (*Server, error) {
	// Setup DB
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Setup Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
	})

	// Gin setup
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestLogger(logger))

	// Routes
	routes.RegisterRoutes(r, db, rdb)


	return &Server{
		engine: r,
		logger: logger,
		cfg:    cfg,
		db:     db,
		redis:  rdb,
	}, nil
}

func (s *Server) Run() error {
	addr := ":8080" // fixed API port
	s.logger.Info("Starting API server", zap.String("addr", addr))
	return s.engine.Run(addr)
}
