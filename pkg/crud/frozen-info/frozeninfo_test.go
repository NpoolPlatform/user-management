package frozeninfo

import (
	"context"
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

func TestFrozenInfoCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	frozenInfo := npool.FrozenUser{
		UserId:      uuid.New().String(),
		FrozenBy:    uuid.New().String(),
		FrozenCause: "use has done some illegal operations",
	}

	// create a frozen user request.
	resp, err := Create(context.Background(), &npool.FrozenUserRequest{
		UserId:      frozenInfo.UserId,
		FrozenBy:    frozenInfo.FrozenBy,
		FrozenCause: frozenInfo.FrozenCause,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.Id, uuid.UUID{})
		assert.Equal(t, resp.Info.UserId, frozenInfo.UserId)
		assert.Equal(t, resp.Info.FrozenBy, frozenInfo.FrozenBy)
		assert.Equal(t, resp.Info.FrozenCause, frozenInfo.FrozenCause)
		assert.Equal(t, resp.Info.Status, FrozenStatus)
		frozenInfo.Id = resp.Info.Id
	}

	// create a frozen user request which is still in frozen.
	_, err = Create(context.Background(), &npool.FrozenUserRequest{
		UserId:      frozenInfo.UserId,
		FrozenBy:    frozenInfo.FrozenBy,
		FrozenCause: frozenInfo.FrozenCause,
	})
	if assert.NotNil(t, err) {
		fmt.Println(err)
	}

	// unfrozen user.
	frozenInfo.UnfrozenBy = uuid.New().String()
	frozenInfo.Id = resp.Info.Id
	resp1, err := Update(context.Background(), &npool.UnfrozenUserRequest{
		Id:         frozenInfo.Id,
		UnfrozenBy: frozenInfo.UnfrozenBy,
		UserId:     frozenInfo.UserId,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.Id, frozenInfo.Id)
		assert.Equal(t, resp1.Info.UserId, frozenInfo.UserId)
		assert.Equal(t, resp1.Info.FrozenBy, frozenInfo.FrozenBy)
		assert.Equal(t, resp1.Info.FrozenCause, frozenInfo.FrozenCause)
		assert.Equal(t, resp1.Info.Status, UnfrozenStatus)
		assert.Equal(t, resp1.Info.UnfrozenBy, frozenInfo.UnfrozenBy)
	}

	// get user frozen list.
	resp2, err := Get(context.Background())
	assert.Nil(t, err)
	fmt.Println("get frozen user list response is:", resp2)
}
