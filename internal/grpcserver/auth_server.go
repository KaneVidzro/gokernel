package grpcserver

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	pb "github.com/kanevidzro/gokernel/proto/auth"
)

type AuthServer struct {
    pb.UnimplementedAuthServiceServer
    JWTSecret []byte
}

func (s *AuthServer) ValidateToken(ctx context.Context, req *pb.TokenRequest) (*pb.TokenClaims, error) {
    token, err := jwt.Parse(req.Token, func(t *jwt.Token) (interface{}, error) {
        return s.JWTSecret, nil
    })
    if err != nil || !token.Valid {
        return &pb.TokenClaims{Valid: false}, nil
    }

    claims := token.Claims.(jwt.MapClaims)
    return &pb.TokenClaims{
        UserId: claims["user_id"].(string),
        Role:   claims["role"].(string),
        Valid:  true,
    }, nil
}
