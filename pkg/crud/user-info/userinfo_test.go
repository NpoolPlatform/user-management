package userinfo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/user-management/message/npool"
	"github.com/NpoolPlatform/user-management/pkg/encryption"
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

func assertUserBasicInfo(t *testing.T, err error, userInfo, myUserInfo *npool.UserBasicInfo) {
	if assert.Nil(t, err) {
		salt, err := GetUserSalt(context.Background(), userInfo.UserId)
		if assert.Nil(t, err) {
			assert.Nil(t, encryption.VerifyUserPassword(myUserInfo.Password, userInfo.Password, salt))
			assert.Equal(t, userInfo.UserId, myUserInfo.UserId)
			assert.Equal(t, userInfo.Username, myUserInfo.Username)
			assert.Equal(t, userInfo.Age, myUserInfo.Age)
			assert.Equal(t, userInfo.Gender, myUserInfo.Gender)
			assert.Equal(t, userInfo.Region, myUserInfo.Region)
			assert.Equal(t, userInfo.Birthday, myUserInfo.Birthday)
			assert.Equal(t, userInfo.Country, myUserInfo.Country)
			assert.Equal(t, userInfo.City, myUserInfo.City)
			assert.Equal(t, userInfo.Province, myUserInfo.Province)
			assert.Equal(t, userInfo.PhoneNumber, myUserInfo.PhoneNumber)
			assert.Equal(t, userInfo.EmailAddress, myUserInfo.EmailAddress)
			assert.Equal(t, userInfo.SignupMethod, myUserInfo.SignupMethod)
			assert.Equal(t, userInfo.Career, myUserInfo.Career)
			assert.Equal(t, userInfo.DisplayName, myUserInfo.DisplayName)
		}
	}
}

func TestUserInfoCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	userInfo := npool.UserBasicInfo{
		UserId:       uuid.New().String(),
		Username:     uuid.New().String(),
		Password:     "123456789",
		Age:          22,
		Gender:       "male",
		Region:       "Asia",
		Birthday:     "1999-06-27",
		Country:      "China",
		Province:     "Shanghai",
		City:         "Shanghai",
		PhoneNumber:  uuid.New().String(),
		EmailAddress: uuid.New().String(),
		SignupMethod: "system",
		Career:       "soft engineer",
		DisplayName:  "Crazyzpl",
	}

	resp, err := Create(context.Background(), &npool.AddUserRequest{
		AppId:    "123456789",
		UserInfo: &userInfo,
	})
	if assert.Nil(t, err) {
		salt, err := GetUserSalt(context.Background(), userInfo.UserId)
		if assert.Nil(t, err) {
			assert.Nil(t, encryption.VerifyUserPassword(userInfo.Password, resp.UserInfo.Password, salt))
			assert.NotEqual(t, resp.UserInfo.UserId, uuid.UUID{})
			assert.Equal(t, resp.UserInfo.Username, userInfo.Username)
			assert.Equal(t, resp.UserInfo.Age, userInfo.Age)
			assert.Equal(t, resp.UserInfo.Gender, userInfo.Gender)
			assert.Equal(t, resp.UserInfo.Region, userInfo.Region)
			assert.Equal(t, resp.UserInfo.Birthday, userInfo.Birthday)
			assert.Equal(t, resp.UserInfo.Country, userInfo.Country)
			assert.Equal(t, resp.UserInfo.City, userInfo.City)
			assert.Equal(t, resp.UserInfo.Province, userInfo.Province)
			assert.Equal(t, resp.UserInfo.PhoneNumber, userInfo.PhoneNumber)
			assert.Equal(t, resp.UserInfo.EmailAddress, userInfo.EmailAddress)
			assert.Equal(t, resp.UserInfo.SignupMethod, userInfo.SignupMethod)
			assert.Equal(t, resp.UserInfo.Career, userInfo.Career)
			assert.Equal(t, resp.UserInfo.DisplayName, userInfo.DisplayName)
		}
	}

	userInfo.DisplayName = "lpzCrazy"
	resp1, err := Update(context.Background(), &npool.UpdateUserInfoRequest{
		UserInfo: &userInfo,
	})
	assertUserBasicInfo(t, err, resp1.UserInfo, &userInfo)

	resp2, err := Get(context.Background(), &npool.GetUserRequest{
		UserId: userInfo.UserId,
	})
	assertUserBasicInfo(t, err, resp2.UserInfo, &userInfo)

	resp4, err := QueryUserByUserID(context.Background(), userInfo.UserId)
	if assert.Nil(t, err) {
		fmt.Println(resp4)
	}

	resp5, err := QueryUserByEmailAddress(context.Background(), userInfo.EmailAddress)
	if assert.Nil(t, err) {
		fmt.Println(resp5)
	}

	resp6, err := QueryUserByPhoneNumber(context.Background(), userInfo.PhoneNumber)
	if assert.Nil(t, err) {
		fmt.Println(resp6)
	}

	resp7, err := QueryUserByUsername(context.Background(), userInfo.Username)
	if assert.Nil(t, err) {
		fmt.Println(resp7)
	}

	resp3, err := Delete(context.Background(), &npool.DeleteUserRequest{
		DeleteUserIds: []string{userInfo.UserId},
	})
	assert.Nil(t, err)
	fmt.Println("delete user response is", resp3)
}
