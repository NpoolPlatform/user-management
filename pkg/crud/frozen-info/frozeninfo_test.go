package frozeninfo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/user"
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

func TestFrozenInfoCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	frozenInfo := npool.FrozenUser{
		UserID:      uuid.New().String(),
		FrozenBy:    uuid.New().String(),
		FrozenCause: "use has done some illegal operations",
	}

	// create a frozen user request.
	resp, err := Create(context.Background(), &npool.FrozenUserRequest{
		UserID:      frozenInfo.UserID,
		FrozenBy:    frozenInfo.FrozenBy,
		FrozenCause: frozenInfo.FrozenCause,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assert.Equal(t, resp.Info.UserID, frozenInfo.UserID)
		assert.Equal(t, resp.Info.FrozenBy, frozenInfo.FrozenBy)
		assert.Equal(t, resp.Info.FrozenCause, frozenInfo.FrozenCause)
		assert.Equal(t, resp.Info.Status, FrozenStatus)
		frozenInfo.ID = resp.Info.ID
	}

	// create a frozen user request which is still in frozen.
	_, err = Create(context.Background(), &npool.FrozenUserRequest{
		UserID:      frozenInfo.UserID,
		FrozenBy:    frozenInfo.FrozenBy,
		FrozenCause: frozenInfo.FrozenCause,
	})
	if assert.NotNil(t, err) {
		fmt.Println(err)
	}

	_, err = Get(context.Background(), &npool.QueryUserFrozenRequest{
		UserID: frozenInfo.UserID,
	})
	assert.Nil(t, err)

	// unfrozen user.
	frozenInfo.UnfrozenBy = uuid.New().String()
	frozenInfo.ID = resp.Info.ID
	resp1, err := Update(context.Background(), &npool.UnfrozenUserRequest{
		ID:         frozenInfo.ID,
		UnfrozenBy: frozenInfo.UnfrozenBy,
		UserID:     frozenInfo.UserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, frozenInfo.ID)
		assert.Equal(t, resp1.Info.UserID, frozenInfo.UserID)
		assert.Equal(t, resp1.Info.FrozenBy, frozenInfo.FrozenBy)
		assert.Equal(t, resp1.Info.FrozenCause, frozenInfo.FrozenCause)
		assert.Equal(t, resp1.Info.Status, UnfrozenStatus)
		assert.Equal(t, resp1.Info.UnfrozenBy, frozenInfo.UnfrozenBy)
	}

	// get user frozen list.
	resp2, err := GetAll(context.Background())
	assert.Nil(t, err)
	fmt.Println("get frozen user list response is:", resp2)
}
