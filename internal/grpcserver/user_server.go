package grpcserver

import (
	"context"

	"github.com/kanevidzro/gokernel/internal/user"
	pb "github.com/kanevidzro/gokernel/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
    pb.UnimplementedUserServiceServer
    Repo *user.Repository
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
    u, err := s.Repo.GetByID(req.Id)
    if err != nil {
        return nil, status.Errorf(codes.NotFound, "user not found")
    }

    return &pb.UserResponse{
        Id:    u.ID,
        Email: u.Email,
        Role:  u.Role,
    }, nil
}
