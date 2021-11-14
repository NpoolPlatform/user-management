package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/user-management/message/npool"
	crud "github.com/NpoolPlatform/user-management/pkg/crud/user-info"
	middleware "github.com/NpoolPlatform/user-management/pkg/middleware/user-info"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SignUp(ctx context.Context, in *npool.SignupRequest) (*npool.SignupResponse, error) {
	resp, err := middleware.Signup(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("user signup error: %v", err)
		return &npool.SignupResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) AddUser(ctx context.Context, in *npool.AddUserRequest) (*npool.AddUserResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("add user error: %v", err)
		return &npool.AddUserResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) GetUser(ctx context.Context, in *npool.GetUserRequest) (*npool.GetUserResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user info error: %v", err)
		return &npool.GetUserResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) GetUsers(ctx context.Context, in *npool.GetUsersRequest) (*npool.GetUsersResponse, error) {
	resp, err := crud.GetAll(ctx)
	if err != nil {
		logger.Sugar().Errorf("get all users info error: %v", err)
		return &npool.GetUsersResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *npool.DeleteUserRequest) (*npool.DeleteUserResponse, error) {
	resp, err := crud.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete user error: %v", err)
		return &npool.DeleteUserResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) UpdateUserInfo(ctx context.Context, in *npool.UpdateUserInfoRequest) (*npool.UpdateUserInfoResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to update user: %v", err)
		return &npool.UpdateUserInfoResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) ChangeUserPassword(ctx context.Context, in *npool.ChangeUserPasswordRequest) (*npool.ChangeUserPasswordResponse, error) {
	resp, err := middleware.ChangeUserPassword(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("failt to change password: %v", err)
		return &npool.ChangeUserPasswordResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) ForgetPassword(ctx context.Context, in *npool.ForgetPasswordRequest) (*npool.ForgetPasswordResponse, error) {
	resp, err := middleware.ForgetPassword(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("forget password error: %v", err)
		return &npool.ForgetPasswordResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) BindUserPhone(ctx context.Context, in *npool.BindUserPhoneRequest) (*npool.BindUserPhoneResponse, error) {
	resp, err := middleware.BindUserPhone(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("bind user phone error: %v", err)
		return &npool.BindUserPhoneResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) BindUserEmail(ctx context.Context, in *npool.BindUserEmailRequest) (*npool.BindUserEmailResponse, error) {
	resp, err := middleware.BindUserEmail(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("bind user email error: %v", err)
		return &npool.BindUserEmailResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) QueryUserExist(ctx context.Context, in *npool.QueryUserExistRequest) (*npool.QueryUserExistResponse, error) {
	resp, err := crud.QueryUserExist(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to query user: %v", err)
		return &npool.QueryUserExistResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}
