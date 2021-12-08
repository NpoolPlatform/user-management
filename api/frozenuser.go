package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/user-management/message/npool"
	crud "github.com/NpoolPlatform/user-management/pkg/crud/frozen-info"
	middleware "github.com/NpoolPlatform/user-management/pkg/middleware/frozen-user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) FrozenUser(ctx context.Context, in *npool.FrozenUserRequest) (*npool.FrozenUserResponse, error) {
	resp, err := middleware.FrozenUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to frozen user: %v", err)
		return &npool.FrozenUserResponse{}, status.Errorf(codes.FailedPrecondition, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) UnfrozenUser(ctx context.Context, in *npool.UnfrozenUserRequest) (*npool.UnfrozenUserResponse, error) {
	resp, err := middleware.UnfrozenUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to unfrozen user: %v", err)
		return &npool.UnfrozenUserResponse{}, status.Errorf(codes.FailedPrecondition, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) QueryUserFrozen(ctx context.Context, in *npool.QueryUserFrozenRequest) (*npool.QueryUserFrozenResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to get user frozen info: %v", err)
		return &npool.QueryUserFrozenResponse{}, status.Errorf(codes.FailedPrecondition, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) GetFrozenUsers(ctx context.Context, in *npool.GetFrozenUsersRequest) (*npool.GetFrozenUsersResponse, error) {
	resp, err := crud.GetAll(ctx)
	if err != nil {
		logger.Sugar().Errorf("fail to get frozen user list: %v", err)
		return &npool.GetFrozenUsersResponse{}, status.Errorf(codes.FailedPrecondition, "internal server error: %v", err)
	}
	return resp, nil
}
