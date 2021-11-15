package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/user-management/message/npool"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func assertUserInfo(t *testing.T, actual, expected *npool.UserBasicInfo) {
	assert.Equal(t, actual.Username, expected.Username)
	assert.Equal(t, actual.EmailAddress, expected.EmailAddress)
	assert.Equal(t, actual.PhoneNumber, expected.PhoneNumber)
}

func TestUserInfoAPI(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	signupUserInfo := npool.UserBasicInfo{
		Username:     "test-signup" + uuid.New().String(),
		Password:     "123456789",
		EmailAddress: uuid.New().String() + ".com",
	}

	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.SignupRequest{
			Username: signupUserInfo.Username,
			Password: signupUserInfo.Password,
		}).
		Post("http://localhost:50070/v1/signup")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp1.StatusCode())
		info := npool.SignupResponse{}
		err := json.Unmarshal(resp1.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.UserId, uuid.UUID{})
			assertUserInfo(t, info.Info, &signupUserInfo)
			signupUserInfo.UserId = info.Info.UserId
		}
	}

	respp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.QueryUserExistRequest{
			Username: signupUserInfo.Username,
			Password: signupUserInfo.Password,
		}).
		Post("http://localhost:50070/v1/query/user/exist")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, respp.StatusCode())
		response := npool.QueryUserExistResponse{}
		err := json.Unmarshal(respp.Body(), &response)
		if assert.Nil(t, err) {
			assert.NotNil(t, response.Info)
		}
	}

	addUserInfo := npool.UserBasicInfo{
		Username:    "test-add" + uuid.New().String(),
		Password:    "123456789",
		PhoneNumber: uuid.New().String(),
	}
	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.AddUserRequest{
			UserInfo: &addUserInfo,
		}).
		Post("http://localhost:50070/v1/add/user")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp2.StatusCode())
		info := npool.AddUserResponse{}
		err := json.Unmarshal(resp2.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.UserId, uuid.UUID{})
			assertUserInfo(t, info.Info, &addUserInfo)
			addUserInfo.UserId = info.Info.UserId
		}
	}

	resp3, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.GetUserRequest{
			UserId: signupUserInfo.UserId,
		}).
		Post("http://localhost:50070/v1/get/user")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp3.StatusCode())
		info := npool.GetUserResponse{}
		err := json.Unmarshal(resp3.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.UserId, uuid.UUID{})
			assertUserInfo(t, info.Info, &signupUserInfo)
		}
	}

	resp4, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.GetUsersRequest{}).
		Post("http://localhost:50070/v1/get/users")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp4.StatusCode())
	}

	resp5, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.UpdateUserInfoRequest{
			Info: &addUserInfo,
		}).
		Post("http://localhost:50070/v1/update/user")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp5.StatusCode())
		info := npool.UpdateUserInfoResponse{}
		err := json.Unmarshal(resp5.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.UserId, uuid.UUID{})
			assertUserInfo(t, info.Info, &addUserInfo)
		}
	}

	resp6, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.ChangeUserPasswordRequest{
			UserId:      signupUserInfo.UserId,
			OldPassword: signupUserInfo.Password,
			Password:    "987654321",
		}).
		Post("http://localhost:50070/v1/change/password")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp6.StatusCode())
	}

	resp7, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.ForgetPasswordRequest{
			PhoneNumber: addUserInfo.PhoneNumber,
			Password:    "123456789",
		}).
		Post("http://localhost:50070/v1/forget/password")
	if assert.Nil(t, err) {
		assert.NotEqual(t, 200, resp7.StatusCode())
	}

	resp8, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.ForgetPasswordRequest{
			EmailAddress: signupUserInfo.EmailAddress,
			Password:     "123456789",
		}).
		Post("http://localhost:50070/v1/forget/password")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp8.StatusCode())
	}

	resp9, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.BindUserPhoneRequest{
			UserId:      addUserInfo.UserId,
			PhoneNumber: uuid.New().String(),
		}).
		Post("http://localhost:50070/v1/bind/phone")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp9.StatusCode())
	}

	resp10, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.BindUserEmailRequest{
			UserId:       addUserInfo.UserId,
			EmailAddress: uuid.New().String(),
		}).
		Post("http://localhost:50070/v1/bind/email")
	if assert.Nil(t, err) {
		assert.NotEqual(t, 200, resp10.StatusCode())
	}

	resp11, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.DeleteUserRequest{
			DeleteUserIds: []string{signupUserInfo.UserId},
		}).
		Post("http://localhost:50070/v1/delete/users")
	fmt.Println(err)
	if assert.Nil(t, err) {
		fmt.Println("delete resp is", resp11)
		assert.Equal(t, 200, resp11.StatusCode())
	}
}
