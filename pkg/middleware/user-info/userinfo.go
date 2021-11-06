package userinfo

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/user-management/message/npool"
	userinfo "github.com/NpoolPlatform/user-management/pkg/crud/user-info"
	"github.com/NpoolPlatform/user-management/pkg/encryption"
	"golang.org/x/xerrors"
)

func Signup(ctx context.Context, in *npool.SignupRequest) (*npool.SignupResponse, error) {
	_, err := userinfo.QueryUserByUsername(ctx, in.Username)
	if err == nil {
		return nil, xerrors.Errorf("user exists")
	}

	request := &npool.AddUserRequest{
		UserInfo: &npool.UserBasicInfo{
			Username:     in.Username,
			Password:     in.Password,
			EmailAddress: in.EmailAddress,
			PhoneNumber:  in.PhoneNumber,
			SignupMethod: "user sign up",
		},
	}

	resp, err := userinfo.Create(ctx, request)
	if err != nil {
		return nil, err
	}

	return &npool.SignupResponse{
		UserInfo: resp.UserInfo,
	}, nil
}

func AddUser(ctx context.Context, in *npool.AddUserRequest) (*npool.AddUserResponse, error) {
	_, err := userinfo.QueryUserByUsername(ctx, in.UserInfo.Username)
	if err == nil {
		return nil, xerrors.Errorf("user exists")
	}

	in.UserInfo.SignupMethod = "admin create"

	resp, err := userinfo.Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return &npool.AddUserResponse{
		UserInfo: resp.UserInfo,
	}, nil
}

func ChangeUserPassword(ctx context.Context, in *npool.ChangeUserPasswordRequest) (*npool.ChangeUserPasswordResponse, error) {
	dbPassword, err := userinfo.GetUserPassword(ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	salt, err := userinfo.GetUserSalt(ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	fmt.Printf("salt is: %v; user id is: %v", salt, in.UserId)
	err = encryption.VerifyUserPassword(in.OldPassword, dbPassword, salt)
	if err != nil {
		return nil, err
	}

	err = userinfo.SetPassword(ctx, in.Password, in.UserId)
	if err != nil {
		return nil, err
	}
	return &npool.ChangeUserPasswordResponse{
		Info: "change user password successfully",
	}, nil
}

func ForgetPassword(ctx context.Context, in *npool.ForgetPasswordRequest) (*npool.ForgetPasswordResponse, error) {
	var userID string
	if in.PhoneNumber == "" && in.EmailAddress != "" {
		userInfo, err := userinfo.QueryUserByEmailAddress(ctx, in.EmailAddress)
		if err != nil {
			return nil, err
		}
		userID = userInfo.UserId
	} else if in.EmailAddress == "" && in.PhoneNumber != "" {
		userInfo, err := userinfo.QueryUserByPhoneNumber(ctx, in.PhoneNumber)
		if err != nil {
			return nil, err
		}
		userID = userInfo.UserId
	}

	err := userinfo.SetPassword(ctx, in.Password, userID)
	if err != nil {
		return nil, err
	}
	return &npool.ForgetPasswordResponse{
		Info: "change password successfully",
	}, nil
}

func BindUserPhone(ctx context.Context, in *npool.BindUserPhoneRequest) (*npool.BindUserPhoneResponse, error) {
	userInfo, err := userinfo.QueryUserByUserID(ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	userInfo.PhoneNumber = in.PhoneNumber

	_, err = userinfo.Update(ctx, &npool.UpdateUserInfoRequest{
		UserInfo: userInfo,
	})
	if err != nil {
		return nil, err
	}
	return &npool.BindUserPhoneResponse{
		Info: "bind phone number successfully",
	}, nil
}

func BindUserEmail(ctx context.Context, in *npool.BindUserEmailRequest) (*npool.BindUserEmailResponse, error) {
	userInfo, err := userinfo.QueryUserByUserID(ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	userInfo.EmailAddress = in.EmailAddress

	_, err = userinfo.Update(ctx, &npool.UpdateUserInfoRequest{
		UserInfo: userInfo,
	})
	if err != nil {
		return nil, err
	}
	return &npool.BindUserEmailResponse{
		Info: "bind email address successfully",
	}, nil
}
