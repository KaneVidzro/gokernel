package main

import (
	"net"

	"github.com/kanevidzro/gokernel/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
    logger.Init()

    grpcPort := "50051" // hardcoded gRPC port

    lis, err := net.Listen("tcp", ":"+grpcPort)
    if err != nil {
        logger.L().Fatal("Failed to listen", zap.Error(err))
    }

    grpcServer := grpc.NewServer()
    // pb.RegisterUserServiceServer(grpcServer, &UserService{})

    logger.L().Info("Starting gRPC server", zap.String("port", grpcPort))
    if err := grpcServer.Serve(lis); err != nil {
        logger.L().Fatal("Failed to serve gRPC", zap.Error(err))
    }
}
