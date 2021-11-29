package userinfo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/user-management/message/npool"
	testinit "github.com/NpoolPlatform/user-management/pkg/test-init"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestUserInfoMiddleware(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	SignupUserInfo := &npool.SignupRequest{
		Username:    "test-signup" + uuid.New().String(),
		Password:    "123456789",
		PhoneNumber: "test-signup" + uuid.New().String(),
		AppID:       "ff2c5d50-be56-413e-aba5-9c7ad888a769",
	}
	CreateUserInfo := &npool.AddUserRequest{
		UserInfo: &npool.UserBasicInfo{
			Username:    "test-add" + uuid.NewString(),
			Password:    "123456789",
			PhoneNumber: "test-add" + uuid.New().String(),
		},
	}
	_, err := Signup(context.Background(), SignupUserInfo)
	assert.NotNil(t, err)

	resp2, err := AddUser(context.Background(), CreateUserInfo)
	if assert.Nil(t, err) {
		fmt.Printf("add user resp is: %v", resp2)
	}

	_, err = ChangeUserPassword(context.Background(), &npool.ChangeUserPasswordRequest{
		UserID:      resp2.Info.UserID,
		OldPassword: SignupUserInfo.Password,
		Password:    "987654321",
	})
	assert.Nil(t, err)

	_, err = ForgetPassword(context.Background(), &npool.ForgetPasswordRequest{
		PhoneNumber: SignupUserInfo.PhoneNumber,
		Password:    "987654321",
	})
	assert.Nil(t, err)

	_, err = BindUserPhone(context.Background(), &npool.BindUserPhoneRequest{
		UserID:      resp2.Info.UserID,
		PhoneNumber: "test-bind" + uuid.New().String(),
	})
	assert.Nil(t, err)

	_, err = BindUserEmail(context.Background(), &npool.BindUserEmailRequest{
		UserID:       resp2.Info.UserID,
		EmailAddress: "test-bind" + uuid.New().String(),
	})
	assert.NotNil(t, err)
}
