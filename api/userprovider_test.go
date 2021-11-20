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

func TestUserProviderAPI(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	addUserInfo := npool.UserBasicInfo{
		Username:    "test-add" + uuid.New().String(),
		Password:    "123456789",
		PhoneNumber: uuid.New().String(),
	}
	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.AddUserRequest{
			UserInfo: &addUserInfo,
		}).
		Post("http://localhost:50070/v1/add/user")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp1.StatusCode())
		info := npool.AddUserResponse{}
		err := json.Unmarshal(resp1.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.UserID, uuid.UUID{})
			assertUserInfo(t, info.Info, &addUserInfo)
			addUserInfo.UserID = info.Info.UserID
		}
	}

	userProvider := npool.UserProvider{
		UserID:         addUserInfo.UserID,
		ProviderID:     uuid.New().String(),
		ProviderUserID: uuid.New().String(),
	}

	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.BindThirdPartyRequest{
			UserID:         userProvider.UserID,
			ProviderID:     userProvider.ProviderID,
			ProviderUserID: userProvider.ProviderUserID,
		}).Post("http://localhost:50070/v1/bind/thirdparty")
	if assert.Nil(t, err) {
		fmt.Println("resp2 is", resp2)
		assert.Equal(t, 200, resp2.StatusCode())
		info := npool.BindThirdPartyResponse{}
		err := json.Unmarshal(resp2.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.ID, uuid.UUID{})
			assert.Equal(t, info.Info.UserID, userProvider.UserID)
			assert.Equal(t, info.Info.ProviderID, userProvider.ProviderID)
			assert.Equal(t, info.Info.ProviderUserID, userProvider.ProviderUserID)
			userProvider.ID = info.Info.ID
		}
	}

	resp5, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.QueryUserByUserProviderIDRequest{
			ProviderID:     userProvider.ProviderID,
			ProviderUserID: userProvider.ProviderUserID,
		}).Post("http://localhost:50070/v1/query/user/by/userproviderid")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp5.StatusCode())
		fmt.Printf("query user by user provider id is: %v", resp5)
	}

	resp3, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.GetUserProvidersRequest{
			UserID: userProvider.UserID,
		}).Post("http://localhost:50070/v1/get/user/providers")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp3.StatusCode())
		fmt.Printf("get user providers list resp is: %v", resp3)
	}

	resp4, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.UnbindThirdPartyRequest{
			UserID:     userProvider.UserID,
			ProviderID: userProvider.ProviderID,
		}).Post("http://localhost:50070/v1/unbind/thirdparty")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp4.StatusCode())
		info := npool.UnbindThirdPartyResponse{}
		err := json.Unmarshal(resp4.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.ID, uuid.UUID{})
			assert.Equal(t, info.Info.UserID, userProvider.UserID)
			assert.Equal(t, info.Info.ProviderID, userProvider.ProviderID)
			fmt.Printf("provider user id is: %v\n", info.Info.ProviderUserID)
		}
	}
}
