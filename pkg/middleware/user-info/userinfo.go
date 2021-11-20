package userinfo

import (
	"context"

	"github.com/NpoolPlatform/user-management/message/npool"
	userinfo "github.com/NpoolPlatform/user-management/pkg/crud/user-info"
	"github.com/NpoolPlatform/user-management/pkg/encryption"
	"github.com/NpoolPlatform/user-management/pkg/grpc"
	"golang.org/x/xerrors"
)

func Signup(ctx context.Context, in *npool.SignupRequest) (*npool.SignupResponse, error) {
	if in.Username != "" {
		_, err := userinfo.QueryUserByUsername(ctx, in.Username)
		if err == nil {
			return nil, xerrors.Errorf("user exists")
		}
	}

	if in.EmailAddress != "" {
		if in.Code == "" {
			return nil, xerrors.Errorf("must have code to verify email")
		}
		_, err := userinfo.QueryUserByUsername(ctx, in.EmailAddress)
		if err == nil {
			return nil, xerrors.Errorf("email has been used")
		}

		err = grpc.VerifyCode(in.EmailAddress, in.Code)
		if err != nil {
			return nil, xerrors.Errorf("input code is wrong")
		}
	}

	if in.PhoneNumber != "" {
		_, err := userinfo.QueryUserByUsername(ctx, in.PhoneNumber)
		if err == nil {
			return nil, xerrors.Errorf("phone number has been used")
		}
	}

	if in.Username == "" {
		in.Username = in.EmailAddress
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

	err = grpc.AddUserToApplication(resp.Info.UserID, in.AppID)
	if err != nil {
		return nil, err
	}

	return &npool.SignupResponse{
		Info: resp.Info,
	}, nil
}

func AddUser(ctx context.Context, in *npool.AddUserRequest) (*npool.AddUserResponse, error) {
	_, err := userinfo.QueryUserByUsername(ctx, in.UserInfo.Username)
	if err == nil {
		return nil, xerrors.Errorf("user exists")
	}

	if in.UserInfo.EmailAddress != "" {
		_, err := userinfo.QueryUserByUsername(ctx, in.UserInfo.EmailAddress)
		if err == nil {
			return nil, xerrors.Errorf("email has been used")
		}
	}

	if in.UserInfo.PhoneNumber != "" {
		_, err := userinfo.QueryUserByUsername(ctx, in.UserInfo.PhoneNumber)
		if err == nil {
			return nil, xerrors.Errorf("phone number has been used")
		}
	}

	in.UserInfo.SignupMethod = "admin create"

	resp, err := userinfo.Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func SetPassword(ctx context.Context, in *npool.SetPasswordRequest) (*npool.SetPasswordResponse, error) {
	resp, err := userinfo.QueryUserByUsername(ctx, in.Username)
	if err != nil {
		return nil, err
	}

	err = userinfo.SetPassword(ctx, in.Password, resp.UserID)
	if err != nil {
		return nil, err
	}

	return &npool.SetPasswordResponse{
		Info: "set password successfully",
	}, nil
}

func ChangeUserPassword(ctx context.Context, in *npool.ChangeUserPasswordRequest) (*npool.ChangeUserPasswordResponse, error) {
	dbPassword, err := userinfo.GetUserPassword(ctx, in.UserID)
	if err != nil {
		return nil, err
	}

	salt, err := userinfo.GetUserSalt(ctx, in.UserID)
	if err != nil {
		return nil, err
	}
	err = encryption.VerifyUserPassword(in.OldPassword, dbPassword, salt)
	if err != nil {
		return nil, err
	}

	err = userinfo.SetPassword(ctx, in.Password, in.UserID)
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
		if in.Code == "" {
			return nil, xerrors.Errorf("input code is empty")
		}
		userInfo, err := userinfo.QueryUserByUsername(ctx, in.EmailAddress)
		if err != nil {
			return nil, err
		}
		err = grpc.VerifyCode(in.EmailAddress, in.Code)
		if err != nil {
			return nil, xerrors.Errorf("fail to verify code: %v", err)
		}
		userID = userInfo.UserID
	} else if in.EmailAddress == "" && in.PhoneNumber != "" {
		userInfo, err := userinfo.QueryUserByUsername(ctx, in.PhoneNumber)
		if err != nil {
			return nil, err
		}
		userID = userInfo.UserID
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
	userInfo, err := userinfo.QueryUserByUserID(ctx, in.UserID)
	if err != nil {
		return nil, err
	}
	userInfo.PhoneNumber = in.PhoneNumber

	_, err = userinfo.Update(ctx, &npool.UpdateUserInfoRequest{
		Info: userInfo,
	})
	if err != nil {
		return nil, err
	}
	return &npool.BindUserPhoneResponse{
		Info: "bind phone number successfully",
	}, nil
}

func BindUserEmail(ctx context.Context, in *npool.BindUserEmailRequest) (*npool.BindUserEmailResponse, error) {
	if in.Code == "" {
		return nil, xerrors.Errorf("input code is empty")
	}

	_, err := userinfo.QueryUserByUsername(ctx, in.EmailAddress)
	if err != xerrors.Errorf("user doesn't exist") {
		return nil, xerrors.Errorf("email has been used: %v", err)
	}

	err = grpc.VerifyCode(in.EmailAddress, in.Code)
	if err != nil {
		return nil, xerrors.Errorf("bind user email error: %v", err)
	}

	userInfo, err := userinfo.QueryUserByUserID(ctx, in.UserID)
	if err != nil {
		return nil, err
	}
	userInfo.EmailAddress = in.EmailAddress

	_, err = userinfo.Update(ctx, &npool.UpdateUserInfoRequest{
		Info: userInfo,
	})
	if err != nil {
		return nil, err
	}
	return &npool.BindUserEmailResponse{
		Info: "bind email address successfully",
	}, nil
}
