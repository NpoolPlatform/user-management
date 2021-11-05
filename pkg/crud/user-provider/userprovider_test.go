package userprovider

import (
	"context"
	"encoding/json"
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

func TestUserProviderCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	userProviderInfo := make(map[string]ProviderUserInfo)
	userProviderInfo["github"] = ProviderUserInfo{
		OauthAvatarURL:   "http://212131.com",
		OauthDisplayName: "crazyzpl",
		OauthEmail:       "1@1.com",
		OauthID:          uuid.New().String(),
		OauthUsername:    "crazyzpl",
	}
	userProviderInfoByte, err := json.Marshal(userProviderInfo)
	assert.Nil(t, err)

	userProvider := npool.UserProvider{
		UserId:           uuid.New().String(),
		ProviderId:       uuid.New().String(),
		ProviderUserId:   userProviderInfo["github"].OauthID,
		UserProviderInfo: string(userProviderInfoByte),
	}
	resp, err := Create(context.Background(), &npool.BindThirdPartyRequest{
		UserId:           userProvider.UserId,
		ProviderId:       userProvider.ProviderId,
		ProviderUserId:   userProvider.ProviderUserId,
		UserProviderInfo: userProvider.UserProviderInfo,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.UserProviderInfo.ID, uuid.UUID{})
		assert.Equal(t, resp.UserProviderInfo.UserId, userProvider.UserId)
		assert.Equal(t, resp.UserProviderInfo.ProviderId, userProvider.ProviderId)
		assert.Equal(t, resp.UserProviderInfo.ProviderUserId, userProvider.ProviderUserId)
		assert.Equal(t, resp.UserProviderInfo.UserProviderInfo, userProvider.UserProviderInfo)
	}

	resp1, err := Get(context.Background(), &npool.GetUserProvidersRequest{
		UserId: userProvider.UserId,
	})
	if assert.Nil(t, err) {
		fmt.Println("get user provider info is", resp1)
	}

	resp2, err := Delete(context.Background(), &npool.UnbindThirdPartyRequest{
		UserId:     userProvider.UserId,
		ProviderId: userProvider.ProviderId,
	})
	if assert.Nil(t, err) {
		fmt.Println("delete user provider info is", resp2)
	}
}
