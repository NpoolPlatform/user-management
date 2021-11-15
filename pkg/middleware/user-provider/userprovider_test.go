package userprovider

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/user-management/message/npool"
	userinfo "github.com/NpoolPlatform/user-management/pkg/crud/user-info"
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

func TestUserProviderMiddleware(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	resp, err := userinfo.Create(context.Background(), &npool.AddUserRequest{
		UserInfo: &npool.UserBasicInfo{
			Username:    "test-frozen" + uuid.New().String(),
			Password:    "123456789",
			PhoneNumber: uuid.New().String(),
		},
	})
	assert.Nil(t, err)

	userProvider := npool.UserProvider{
		UserId:         resp.Info.UserId,
		ProviderId:     uuid.New().String(),
		ProviderUserId: "test-provider" + uuid.New().String(),
	}

	resp1, err := BindThirdParty(context.Background(), &npool.BindThirdPartyRequest{
		UserId:         userProvider.UserId,
		ProviderId:     userProvider.ProviderId,
		ProviderUserId: userProvider.ProviderUserId,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, uuid.UUID{}, resp1.Info.ID)
		assert.Equal(t, userProvider.UserId, resp1.Info.UserId)
		assert.Equal(t, userProvider.ProviderId, resp1.Info.ProviderId)
		assert.Equal(t, userProvider.ProviderUserId, resp1.Info.ProviderUserId)
		userProvider.ID = resp1.Info.ID
	}

	resp4, err := QueryUserByUserProviderID(context.Background(), &npool.QueryUserByUserProviderIDRequest{
		ProviderID:     userProvider.ProviderId,
		ProviderUserID: userProvider.ProviderUserId,
	})
	if assert.Nil(t, err) {
		if assert.NotNil(t, resp4) {
			assert.NotNil(t, resp4.Info.UserBasicInfo)
			assert.NotNil(t, resp4.Info.UserProviderInfo)
		}
	}

	resp2, err := GetUserProviders(context.Background(), &npool.GetUserProvidersRequest{
		UserId: userProvider.UserId,
	})
	assert.Nil(t, err)
	fmt.Printf("get user providers list is: %v\n", resp2)

	resp3, err := UnbindUserProviders(context.Background(), &npool.UnbindThirdPartyRequest{
		UserId:     userProvider.UserId,
		ProviderId: userProvider.ProviderId,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, userProvider.ID)
		assert.Equal(t, resp3.Info.UserId, userProvider.UserId)
		assert.Equal(t, resp3.Info.ProviderId, userProvider.ProviderId)
		fmt.Printf("provider user id is: %v\n", resp3.Info.ProviderUserId)
	}
}
