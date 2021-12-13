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
			Username:    uuid.New().String()[0:12],
			Password:    "123456789",
			PhoneNumber: uuid.New().String(),
		},
	})
	assert.Nil(t, err)

	userProvider := npool.UserProvider{
		UserID:         resp.Info.UserID,
		ProviderID:     uuid.New().String(),
		ProviderUserID: "test-provider" + uuid.New().String(),
	}

	resp1, err := BindThirdParty(context.Background(), &npool.BindThirdPartyRequest{
		UserID:         userProvider.UserID,
		ProviderID:     userProvider.ProviderID,
		ProviderUserID: userProvider.ProviderUserID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, uuid.UUID{}, resp1.Info.ID)
		assert.Equal(t, userProvider.UserID, resp1.Info.UserID)
		assert.Equal(t, userProvider.ProviderID, resp1.Info.ProviderID)
		assert.Equal(t, userProvider.ProviderUserID, resp1.Info.ProviderUserID)
		userProvider.ID = resp1.Info.ID
	}

	resp4, err := QueryUserByUserProviderID(context.Background(), &npool.QueryUserByUserProviderIDRequest{
		ProviderID:     userProvider.ProviderID,
		ProviderUserID: userProvider.ProviderUserID,
	})
	if assert.Nil(t, err) {
		if assert.NotNil(t, resp4) {
			assert.NotNil(t, resp4.Info.UserBasicInfo)
			assert.NotNil(t, resp4.Info.UserProviderInfo)
		}
	}

	resp2, err := GetUserProviders(context.Background(), &npool.GetUserProvidersRequest{
		UserID: userProvider.UserID,
	})
	assert.Nil(t, err)
	fmt.Printf("get user providers list is: %v\n", resp2)

	resp3, err := UnbindUserProviders(context.Background(), &npool.UnbindThirdPartyRequest{
		UserID:     userProvider.UserID,
		ProviderID: userProvider.ProviderID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, userProvider.ID)
		assert.Equal(t, resp3.Info.UserID, userProvider.UserID)
		assert.Equal(t, resp3.Info.ProviderID, userProvider.ProviderID)
		fmt.Printf("provider user id is: %v\n", resp3.Info.ProviderUserID)
	}
}
