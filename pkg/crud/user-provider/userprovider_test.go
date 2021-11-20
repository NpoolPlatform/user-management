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
		UserID:           uuid.New().String(),
		ProviderID:       uuid.New().String(),
		ProviderUserID:   userProviderInfo["github"].OauthID,
		UserProviderInfo: string(userProviderInfoByte),
	}
	resp, err := Create(context.Background(), &npool.BindThirdPartyRequest{
		UserID:           userProvider.UserID,
		ProviderID:       userProvider.ProviderID,
		ProviderUserID:   userProvider.ProviderUserID,
		UserProviderInfo: userProvider.UserProviderInfo,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assert.Equal(t, resp.Info.UserID, userProvider.UserID)
		assert.Equal(t, resp.Info.ProviderID, userProvider.ProviderID)
		assert.Equal(t, resp.Info.ProviderUserID, userProvider.ProviderUserID)
		assert.Equal(t, resp.Info.UserProviderInfo, userProvider.UserProviderInfo)
	}

	resp3, err := QueryUserProviderInfoByProviderUserID(context.Background(), &npool.QueryUserByUserProviderIDRequest{
		ProviderID:     userProvider.ProviderID,
		ProviderUserID: userProvider.ProviderUserID,
	})
	if assert.Nil(t, err) {
		if assert.NotNil(t, resp3) {
			assert.NotNil(t, resp3.Info.UserProviderInfo)
		}
	}

	resp1, err := Get(context.Background(), &npool.GetUserProvidersRequest{
		UserID: userProvider.UserID,
	})
	if assert.Nil(t, err) {
		fmt.Println("get user provider info is", resp1)
	}

	resp2, err := Delete(context.Background(), &npool.UnbindThirdPartyRequest{
		UserID:     userProvider.UserID,
		ProviderID: userProvider.ProviderID,
	})
	if assert.Nil(t, err) {
		fmt.Println("delete user provider info is", resp2)
	}
}
