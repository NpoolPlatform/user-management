package userinfo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/user"
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
		if assert.Nil(t, err) {
			assert.Equal(t, userInfo.UserID, myUserInfo.UserID)
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
		Username:     uuid.New().String()[0:12],
		Password:     "12345dasdas6789",
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
		AppID:    "123456789",
		UserInfo: &userInfo,
	})
	if assert.Nil(t, err) {
		salt, err := GetUserSalt(context.Background(), resp.Info.UserID)
		if assert.Nil(t, err) {
			userInfo.UserID = resp.Info.UserID
			assert.Nil(t, encryption.VerifyUserPassword(userInfo.Password, resp.Info.Password, salt))
			assert.NotEqual(t, resp.Info.UserID, uuid.UUID{})
			assert.Equal(t, resp.Info.Username, userInfo.Username)
			assert.Equal(t, resp.Info.Age, userInfo.Age)
			assert.Equal(t, resp.Info.Gender, userInfo.Gender)
			assert.Equal(t, resp.Info.Region, userInfo.Region)
			assert.Equal(t, resp.Info.Birthday, userInfo.Birthday)
			assert.Equal(t, resp.Info.Country, userInfo.Country)
			assert.Equal(t, resp.Info.City, userInfo.City)
			assert.Equal(t, resp.Info.Province, userInfo.Province)
			assert.Equal(t, resp.Info.PhoneNumber, userInfo.PhoneNumber)
			assert.Equal(t, resp.Info.EmailAddress, userInfo.EmailAddress)
			assert.Equal(t, resp.Info.SignupMethod, userInfo.SignupMethod)
			assert.Equal(t, resp.Info.Career, userInfo.Career)
			assert.Equal(t, resp.Info.DisplayName, userInfo.DisplayName)
		}
	}

	resp10, err := QueryUserExist(context.Background(), &npool.QueryUserExistRequest{
		Username: userInfo.Username,
		Password: userInfo.Password,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp10)
	}

	resp12, err := QueryUserExist(context.Background(), &npool.QueryUserExistRequest{
		Username: userInfo.EmailAddress,
		Password: userInfo.Password,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp12)
	}

	userInfo.DisplayName = "lpzCrazy"
	fmt.Println(userInfo.UserID)
	resp1, err := Update(context.Background(), &npool.UpdateUserInfoRequest{
		Info: &userInfo,
	})
	assertUserBasicInfo(t, err, resp1.Info, &userInfo)

	err = SetPassword(context.Background(), userInfo.Password, userInfo.UserID)
	assert.Nil(t, err)

	resp2, err := Get(context.Background(), &npool.GetUserRequest{
		UserID: userInfo.UserID,
	})
	assertUserBasicInfo(t, err, resp2.Info, &userInfo)

	resp4, err := QueryUserByUserID(context.Background(), userInfo.UserID)
	if assert.Nil(t, err) {
		fmt.Println(resp4)
	}

	resp7, err := QueryUserByParam(context.Background(), userInfo.Username)
	if assert.Nil(t, err) {
		fmt.Println(resp7)
	}

	resp8, err := GetAll(context.Background())
	if assert.Nil(t, err) {
		fmt.Printf("get all user is: %v\n", resp8)
	}

	resp9, err := GetUserPassword(context.Background(), userInfo.UserID)
	if assert.Nil(t, err) {
		fmt.Printf("get user password is: %v\n", resp9)
	}

	resp3, err := Delete(context.Background(), &npool.DeleteUserRequest{
		DeleteUserIDs: []string{userInfo.UserID},
	})
	assert.Nil(t, err)
	fmt.Println("delete user response is", resp3)
}
