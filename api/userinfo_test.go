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

	addUserInfo := npool.UserBasicInfo{
		Username:    uuid.New().String()[0:12],
		Password:    "123456789",
		PhoneNumber: uuid.New().String()[0:12],
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
			assert.NotEqual(t, info.Info.UserID, uuid.UUID{})
			assertUserInfo(t, info.Info, &addUserInfo)
			addUserInfo.UserID = info.Info.UserID
		}
	}

	signupUserInfo := npool.UserBasicInfo{
		Username:     uuid.New().String()[0:12],
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
		assert.NotEqual(t, 200, resp1.StatusCode())
	}

	respp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.QueryUserExistRequest{
			Username: addUserInfo.Username,
			Password: addUserInfo.Password,
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

	resp3, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.GetUserRequest{
			UserID: addUserInfo.UserID,
		}).
		Post("http://localhost:50070/v1/get/user")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp3.StatusCode())
		info := npool.GetUserResponse{}
		err := json.Unmarshal(resp3.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.UserID, uuid.UUID{})
			assertUserInfo(t, info.Info, &addUserInfo)
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
		fmt.Println("add user info user id is:", addUserInfo.UserID)
		fmt.Println("user info api test resp5 is:", resp5)
		assert.Equal(t, 200, resp5.StatusCode())
		info := npool.UpdateUserInfoResponse{}
		err := json.Unmarshal(resp5.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.UserID, uuid.UUID{})
			assertUserInfo(t, info.Info, &addUserInfo)
		}
	}

	resp6, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.ChangeUserPasswordRequest{
			UserID:      addUserInfo.UserID,
			OldPassword: addUserInfo.Password,
			Password:    "987654321",
		}).
		Post("http://localhost:50070/v1/change/password")
	if assert.Nil(t, err) {
		assert.NotEqual(t, 200, resp6.StatusCode())
	}

	resp7, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.ForgetPasswordRequest{
			VerifyParam: addUserInfo.PhoneNumber,
			Password:    "123456789",
		}).
		Post("http://localhost:50070/v1/forget/password")
	if assert.Nil(t, err) {
		assert.NotEqual(t, 200, resp7.StatusCode())
	}

	resp8, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.ForgetPasswordRequest{
			VerifyParam: addUserInfo.EmailAddress,
			Password:    "123456789",
		}).
		Post("http://localhost:50070/v1/forget/password")
	if assert.Nil(t, err) {
		assert.NotEqual(t, 200, resp8.StatusCode())
	}

	resp9, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.BindUserPhoneRequest{
			UserID:      addUserInfo.UserID,
			PhoneNumber: uuid.New().String(),
		}).
		Post("http://localhost:50070/v1/bind/phone")
	if assert.Nil(t, err) {
		assert.NotEqual(t, 200, resp9.StatusCode())
	}

	resp10, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.BindUserEmailRequest{
			UserID:       addUserInfo.UserID,
			EmailAddress: uuid.New().String(),
		}).
		Post("http://localhost:50070/v1/bind/email")
	if assert.Nil(t, err) {
		assert.NotEqual(t, 200, resp10.StatusCode())
	}

	resp11, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.DeleteUserRequest{
			DeleteUserIDs: []string{addUserInfo.UserID},
		}).
		Post("http://localhost:50070/v1/delete/users")
	fmt.Println(err)
	if assert.Nil(t, err) {
		fmt.Println("delete resp is", resp11)
		assert.Equal(t, 200, resp11.StatusCode())
	}
}
