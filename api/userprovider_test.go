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

	createUser := npool.UserBasicInfo{
		Username:     "test-provider" + uuid.New().String(),
		Password:     "123456789",
		EmailAddress: uuid.New().String() + ".com",
	}
	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.SignupRequest{
			Username:     createUser.Username,
			Password:     createUser.Password,
			EmailAddress: createUser.EmailAddress,
		}).
		Post("http://localhost:50070/v1/signup")
	fmt.Println("sign up error", err)
	if assert.Nil(t, err) {
		fmt.Println("resp1 is", resp1)
		assert.Equal(t, 200, resp1.StatusCode())
		info := npool.SignupResponse{}
		err := json.Unmarshal(resp1.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.UserId, uuid.UUID{})
			assertUserInfo(t, info.Info, &createUser)
			createUser.UserId = info.Info.UserId
		}
	}

	userProvider := npool.UserProvider{
		UserId:         createUser.UserId,
		ProviderId:     uuid.New().String(),
		ProviderUserId: uuid.New().String(),
	}

	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.BindThirdPartyRequest{
			UserId:         userProvider.UserId,
			ProviderId:     userProvider.ProviderId,
			ProviderUserId: userProvider.ProviderUserId,
		}).Post("http://localhost:50070/v1/bind/thirdparty")
	if assert.Nil(t, err) {
		fmt.Println("resp2 is", resp2)
		assert.Equal(t, 200, resp2.StatusCode())
		info := npool.BindThirdPartyResponse{}
		err := json.Unmarshal(resp2.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.ID, uuid.UUID{})
			assert.Equal(t, info.Info.UserId, userProvider.UserId)
			assert.Equal(t, info.Info.ProviderId, userProvider.ProviderId)
			assert.Equal(t, info.Info.ProviderUserId, userProvider.ProviderUserId)
			userProvider.ID = info.Info.ID
		}
	}

	resp5, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.QueryUserByUserProviderIDRequest{
			ProviderID:     userProvider.ProviderId,
			ProviderUserID: userProvider.ProviderUserId,
		}).Post("http://localhost:50070/v1/query/user/by/userproviderid")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp5.StatusCode())
		fmt.Printf("query user by user provider id is: %v", resp5)
	}

	resp3, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.GetUserProvidersRequest{
			UserId: userProvider.UserId,
		}).Post("http://localhost:50070/v1/get/user/providers")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp3.StatusCode())
		fmt.Printf("get user providers list resp is: %v", resp3)
	}

	resp4, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.UnbindThirdPartyRequest{
			UserId:     userProvider.UserId,
			ProviderId: userProvider.ProviderId,
		}).Post("http://localhost:50070/v1/unbind/thirdparty")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp4.StatusCode())
		info := npool.UnbindThirdPartyResponse{}
		err := json.Unmarshal(resp4.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.ID, uuid.UUID{})
			assert.Equal(t, info.Info.UserId, userProvider.UserId)
			assert.Equal(t, info.Info.ProviderId, userProvider.ProviderId)
			fmt.Printf("provider user id is: %v\n", info.Info.ProviderUserId)
		}
	}
}
