package frozenuser

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

func TestFrozenUserMiddleware(t *testing.T) { // nolint
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

	frozenUserInfo := npool.FrozenUser{
		UserID:      resp.Info.UserID,
		FrozenBy:    uuid.New().String(),
		FrozenCause: "user has some illegal operations",
	}

	resp1, err := FrozenUser(context.Background(), &npool.FrozenUserRequest{
		UserID:      frozenUserInfo.UserID,
		FrozenBy:    frozenUserInfo.FrozenBy,
		FrozenCause: frozenUserInfo.FrozenCause,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp1.Info.ID, uuid.UUID{})
		assert.Equal(t, resp1.Info.UserID, frozenUserInfo.UserID)
		assert.Equal(t, resp1.Info.FrozenBy, frozenUserInfo.FrozenBy)
		assert.Equal(t, resp1.Info.FrozenCause, frozenUserInfo.FrozenCause)
		frozenUserInfo.ID = resp1.Info.ID
	}

	frozenUserInfo.UnfrozenBy = uuid.New().String()
	resp2, err := UnfrozenUser(context.Background(), &npool.UnfrozenUserRequest{
		ID:         frozenUserInfo.ID,
		UserID:     frozenUserInfo.UserID,
		UnfrozenBy: frozenUserInfo.UnfrozenBy,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, frozenUserInfo.ID)
		assert.Equal(t, resp2.Info.UserID, frozenUserInfo.UserID)
		assert.Equal(t, resp2.Info.FrozenBy, frozenUserInfo.FrozenBy)
		assert.Equal(t, resp2.Info.FrozenCause, frozenUserInfo.FrozenCause)
		assert.Equal(t, resp2.Info.UnfrozenBy, frozenUserInfo.UnfrozenBy)
	}
}
