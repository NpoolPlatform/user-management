package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/user-management/message/npool"
	crud "github.com/NpoolPlatform/user-management/pkg/crud/frozen-info"
	middleware "github.com/NpoolPlatform/user-management/pkg/middleware/frozen-user"
)

func (s *Server) FrozenUser(ctx context.Context, in *npool.FrozenUserRequest) (*npool.FrozenUserResponse, error) {
	resp, err := middleware.FrozenUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to frozen user: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *Server) UnfrozenUser(ctx context.Context, in *npool.UnfrozenUserRequest) (*npool.UnfrozenUserResponse, error) {
	resp, err := middleware.UnfrozenUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to unfrozen user: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *Server) GetFrozenUsers(ctx context.Context, in *npool.GetFrozenUsersRequest) (*npool.GetFrozenUsersResponse, error) {
	resp, err := crud.Get(ctx)
	if err != nil {
		logger.Sugar().Errorf("fail to get frozen user list: %v", err)
		return nil, err
	}
	return resp, nil
}
