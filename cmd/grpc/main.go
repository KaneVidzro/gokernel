package main

import (
	"net"

	"github.com/kanevidzro/gokernel/internal/grpcserver"
	"github.com/kanevidzro/gokernel/pkg/config"
	"github.com/kanevidzro/gokernel/pkg/logger"
	pb "github.com/kanevidzro/gokernel/proto/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
    logger.Init()
    cfg := config.Load() // Load your config

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        logger.L().Fatal("Failed to listen", zap.Error(err))
    }

    grpcServer := grpc.NewServer()

    authServer := &grpcserver.AuthServer{JWTSecret: []byte(cfg.JWTSecret)}
    pb.RegisterAuthServiceServer(grpcServer, authServer)

    logger.L().Info("Starting gRPC server on port 50051")
    if err := grpcServer.Serve(lis); err != nil {
        logger.L().Fatal("Failed to serve gRPC", zap.Error(err))
    }
}
