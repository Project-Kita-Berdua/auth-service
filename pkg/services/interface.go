package services

import (
	"context"

	"github.com/storyofhis/go-grpc-auth-svc/pkg/pb"
)

type AuthSvcInterface interface {
	Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}
