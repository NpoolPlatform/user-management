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

const (
	Email  = "email"
	Phone  = "phone"
	Google = "google"
)

func Signup(ctx context.Context, in *npool.SignupRequest) (*npool.SignupResponse, error) { // nolint
	if in.Code == "" || in.Password == "" {
		return nil, xerrors.Errorf("verify code and password can not empty")
	}

	if in.EmailAddress == "" && in.PhoneNumber == "" {
		return nil, xerrors.Errorf("user account info cannot be null")
	}

	if match := utils.RegexpPassword(in.Password); !match {
		return nil, xerrors.Errorf("password not legal")
	}

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
	if !utils.RegexpPassword(in.UserInfo.Password) {
		return nil, xerrors.Errorf("password isn't legal")
	}

	if in.UserInfo.Username == "" {
		username, err := utils.GenerateUsername()
		if err != nil {
			return nil, xerrors.Errorf("fail to generate username: %v", err)
		}
		in.UserInfo.Username = username
	} else {
		ok := utils.RegexpUsername(in.UserInfo.Username)
		if !ok {
			return nil, xerrors.Errorf("username not legal")
		}
	}

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
	if in.Code == "" || in.Password == "" {
		return nil, xerrors.Errorf("input code or password cannot empty")
	}

	if match := utils.RegexpPassword(in.Password); !match {
		return nil, xerrors.Errorf("password isn't legal")
	}

	if in.VerifyType == Email || in.VerifyType == Phone {
		err := grpc.VerifyCodeWithUserID(in.UserID, in.VerifyParam, in.Code)
		if err != nil {
			return nil, xerrors.Errorf("fail to verify code: %v", err)
		}
	} else {
		return nil, xerrors.Errorf("please input correct user account info")
	}

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
	if in.Code == "" || in.Password == "" {
		return nil, xerrors.Errorf("input code or password cannot be empty")
	}

	if match := utils.RegexpPassword(in.Password); !match {
		return nil, xerrors.Errorf("password isn't legal")
	}

	var userID string
	if in.VerifyType == Email || in.VerifyType == Phone {
		userInfo, err := userinfo.QueryUserByParam(ctx, in.VerifyParam)
		if err != nil {
			return nil, err
		}
		err = grpc.VerifyCode(in.VerifyParam, in.Code)
		if err != nil {
			return nil, xerrors.Errorf("fail to verify code: %v", err)
		}
		userID = userInfo.UserID
	} else {
		return nil, xerrors.Errorf("please input correct user account info")
	}

	err := userinfo.SetPassword(ctx, in.Password, userID)
	if err != nil {
		return nil, err
	}
	return &npool.ForgetPasswordResponse{
		Info: "reset password successfully",
	}, nil
}

func BindUserPhone(ctx context.Context, in *npool.BindUserPhoneRequest) (*npool.BindUserPhoneResponse, error) {
	if in.Code == "" || in.PhoneNumber == "" {
		return nil, xerrors.Errorf("input phone number and code cannot be empty")
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
	if in.Code == "" || in.EmailAddress == "" {
		return nil, xerrors.Errorf("input email address and code cannot be empty")
	}

	err := grpc.VerifyCode(in.EmailAddress, in.Code)
	if err != nil {
		return nil, xerrors.Errorf("bind user email error: %v", err)
	}

	err = userinfo.SetUserEmail(ctx, in.UserID, in.EmailAddress)
	if err != nil {
		return nil, err
	}

	return &npool.BindUserEmailResponse{
		Info: "bind email address successfully",
	}, nil
}

func UpdateUserEmail(ctx context.Context, in *npool.UpdateUserEmailRequest) (*npool.UpdateUserEmailResponse, error) { // nolint
	if in.OldCode == "" || in.NewCode == "" {
		return nil, xerrors.Errorf("input code cannot be empty")
	}

	if in.OldEmail == "" || in.NewEmail == "" {
		return nil, xerrors.Errorf("input old email and new email cannot be empty")
	}

	if in.OldEmail == in.NewEmail {
		return nil, xerrors.Errorf("old email and new email cannot be same")
	}

	err := grpc.VerifyCodeWithUserID(in.UserID, in.OldEmail, in.OldCode)
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

func UpdateUserPhone(ctx context.Context, in *npool.UpdateUserPhoneRequest) (*npool.UpdateUserPhoneResponse, error) { // nolint
	if in.OldCode == "" || in.NewCode == "" {
		return nil, xerrors.Errorf("input code cannot be empty")
	}

	if in.OldPhone == "" || in.NewPhone == "" {
		return nil, xerrors.Errorf("input old email and new email cannot be empty")
	}

	if in.OldPhone == in.NewPhone {
		return nil, xerrors.Errorf("old phone and new phone cannot be same")
	}

	err := grpc.VerifyCodeWithUserID(in.UserID, in.OldPhone, in.OldCode)
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
