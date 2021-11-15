package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/user-management/message/npool"
	middleware "github.com/NpoolPlatform/user-management/pkg/middleware/user-provider"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) BindThirdParty(ctx context.Context, in *npool.BindThirdPartyRequest) (*npool.BindThirdPartyResponse, error) {
	resp, err := middleware.BindThirdParty(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to bind third party: %v", err)
		return &npool.BindThirdPartyResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) UnbindThirdParty(ctx context.Context, in *npool.UnbindThirdPartyRequest) (*npool.UnbindThirdPartyResponse, error) {
	resp, err := middleware.UnbindUserProviders(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to unbind third party: %v", err)
		return &npool.UnbindThirdPartyResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) GetUserProviders(ctx context.Context, in *npool.GetUserProvidersRequest) (*npool.GetUserProvidersResponse, error) {
	resp, err := middleware.GetUserProviders(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to get user providers third party: %v", err)
		return &npool.GetUserProvidersResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) QueryUserByUserProviderID(ctx context.Context, in *npool.QueryUserByUserProviderIDRequest) (*npool.QueryUserByUserProviderIDResponse, error) {
	resp, err := middleware.QueryUserByUserProviderID(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to get user by user provider id: %v", err)
		return &npool.QueryUserByUserProviderIDResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}
