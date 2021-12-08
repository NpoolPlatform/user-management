package userinfo

import (
	"context"

	"github.com/NpoolPlatform/user-management/message/npool"
	userinfo "github.com/NpoolPlatform/user-management/pkg/crud/user-info"
	"github.com/NpoolPlatform/user-management/pkg/encryption"
	"github.com/NpoolPlatform/user-management/pkg/grpc"
	"github.com/NpoolPlatform/user-management/pkg/utils"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func Signup(ctx context.Context, in *npool.SignupRequest) (*npool.SignupResponse, error) {
	err := grpc.QueryAppExist(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("user sign up app not exist: %v", err)
	}

	if in.EmailAddress != "" {
		if in.Code == "" {
			return nil, xerrors.Errorf("must have code to verify email")
		}
		_, err := userinfo.QueryUserByParam(ctx, in.EmailAddress)
		if err == nil {
			return nil, xerrors.Errorf("email has been used")
		}

		err = grpc.VerifyCode(in.EmailAddress, in.Code)
		if err != nil {
			return nil, xerrors.Errorf("input code is wrong")
		}
	}

	if in.PhoneNumber != "" {
		if in.Code == "" {
			return nil, xerrors.Errorf("must have code to verify email")
		}
		_, err := userinfo.QueryUserByParam(ctx, in.PhoneNumber)
		if err == nil {
			return nil, xerrors.Errorf("phone number has been used")
		}

		err = grpc.VerifyCode(in.PhoneNumber, in.Code)
		if err != nil {
			return nil, xerrors.Errorf("input code is wrong")
		}
	}

	username, err := utils.GenerateUsername()
	if err != nil {
		return nil, xerrors.Errorf("fail to generate username: %v", err)
	}

	request := &npool.AddUserRequest{
		UserInfo: &npool.UserBasicInfo{
			Username:     username,
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
	_, err := userinfo.QueryUserByParam(ctx, in.UserInfo.Username)
	if err == nil {
		return nil, xerrors.Errorf("user exists")
	}

	if in.UserInfo.EmailAddress != "" {
		_, err := userinfo.QueryUserByParam(ctx, in.UserInfo.EmailAddress)
		if err == nil {
			return nil, xerrors.Errorf("email has been used")
		}
	}

	if in.UserInfo.PhoneNumber != "" {
		_, err := userinfo.QueryUserByParam(ctx, in.UserInfo.PhoneNumber)
		if err == nil {
			return nil, xerrors.Errorf("phone number has been used")
		}
	}

	in.UserInfo.SignupMethod = "admin create"
	if in.UserInfo.Username == "" {
		username, err := utils.GenerateUsername()
		if err != nil {
			return nil, xerrors.Errorf("fail to generate username: %v", err)
		}
		in.UserInfo.Username = username
	}

	resp, err := userinfo.Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func SetPassword(ctx context.Context, in *npool.SetPasswordRequest) (*npool.SetPasswordResponse, error) {
	resp, err := userinfo.QueryUserByParam(ctx, in.Username)
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

	resp, err := userinfo.QueryUserByUserID(ctx, in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("fail to query user: %v", err)
	}

	err = grpc.VerifyCode(resp.EmailAddress, in.EmailVerifyCode)
	if err != nil {
		return nil, err
	}

	err = grpc.VerifyGoogleCode(in.UserID, in.AppID, in.GoogleVerifyCode)
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
		userInfo, err := userinfo.QueryUserByParam(ctx, in.EmailAddress)
		if err != nil {
			return nil, err
		}
		err = grpc.VerifyCode(in.EmailAddress, in.Code)
		if err != nil {
			return nil, xerrors.Errorf("fail to verify code: %v", err)
		}
		userID = userInfo.UserID
	} else if in.EmailAddress == "" && in.PhoneNumber != "" {
		userInfo, err := userinfo.QueryUserByParam(ctx, in.PhoneNumber)
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
	if in.Code == "" {
		return nil, xerrors.Errorf("input code is empty")
	}

	err := grpc.VerifyCode(in.PhoneNumber, in.Code)
	if err != nil {
		return nil, xerrors.Errorf("fail to verify phone code: %v", err)
	}

	err = userinfo.SetUserPhone(ctx, in.UserID, in.PhoneNumber)
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

	err := grpc.VerifyCode(in.EmailAddress, in.Code)
	if err != nil {
		return nil, xerrors.Errorf("bind user email error: %v", err)
	}

	err = userinfo.SetUserPhone(ctx, in.UserID, in.EmailAddress)
	if err != nil {
		return nil, err
	}

	return &npool.BindUserEmailResponse{
		Info: "bind email address successfully",
	}, nil
}

func UpdateUserEmail(ctx context.Context, in *npool.UpdateUserEmailRequest) (*npool.UpdateUserEmailResponse, error) {
	if in.OldCode == "" || in.NewCode == "" {
		return nil, xerrors.Errorf("input code cannot be empty")
	}

	if in.OldEmail != in.NewEmail {
		return nil, xerrors.Errorf("old email and new email cannot be same")
	}

	err := grpc.VerifyCode(in.OldEmail, in.OldCode)
	if err != nil {
		return nil, xerrors.Errorf("fail to verify old email code: %v", err)
	}

	err = grpc.VerifyCode(in.NewEmail, in.NewCode)
	if err != nil {
		return nil, xerrors.Errorf("fail to verify new email code: %v", err)
	}

	err = userinfo.SetUserEmail(ctx, in.UserID, in.NewEmail)
	if err != nil {
		return nil, xerrors.Errorf("fail to update user email: %v", err)
	}

	return &npool.UpdateUserEmailResponse{
		Info: "Update email successfully",
	}, nil
}

func UpdateUserPhone(ctx context.Context, in *npool.UpdateUserPhoneRequest) (*npool.UpdateUserPhoneResponse, error) {
	if in.OldCode == "" || in.NewCode == "" {
		return nil, xerrors.Errorf("input code cannot be empty")
	}

	if in.OldPhone != in.NewPhone {
		return nil, xerrors.Errorf("old phone and new phone cannot be same")
	}

	err := grpc.VerifyCode(in.OldPhone, in.OldCode)
	if err != nil {
		return nil, xerrors.Errorf("fail to verify old email code: %v", err)
	}

	err = grpc.VerifyCode(in.NewPhone, in.NewCode)
	if err != nil {
		return nil, xerrors.Errorf("fail to verify new email code: %v", err)
	}

	err = userinfo.SetUserPhone(ctx, in.UserID, in.NewPhone)
	if err != nil {
		return nil, xerrors.Errorf("fail to update user phone: %v", err)
	}

	return &npool.UpdateUserPhoneResponse{
		Info: "Update phone successfully",
	}, nil
}

func GetUserDetails(ctx context.Context, in *npool.GetUserDetailsRequest) (*npool.GetUserDetailsResponse, error) {
	if _, err := uuid.Parse(in.UserID); err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	if _, err := uuid.Parse(in.AppID); err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	resp, err := grpc.GetUserApplicationInfo(in.UserID, in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("grpc applciation error: %v", err)
	}

	respUser, err := userinfo.Get(ctx, &npool.GetUserRequest{
		UserID: in.UserID,
	})
	if err != nil {
		return nil, err
	}
	return &npool.GetUserDetailsResponse{
		Info: &npool.UserDetails{
			UserBasicInfo: respUser.Info,
			UserAppInfo:   resp,
		},
	}, nil
}
