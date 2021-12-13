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

func TestFrozenUserAPI(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	addUserInfo := npool.UserBasicInfo{
		Username:    uuid.New().String()[0:12],
		Password:    "12dasda3456789",
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

	frozenUserInfo := npool.FrozenUser{
		UserID:      addUserInfo.UserID,
		FrozenBy:    uuid.New().String(),
		FrozenCause: "user has done some illegal operations",
	}

	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.FrozenUserRequest{
			UserID:      frozenUserInfo.UserID,
			FrozenBy:    frozenUserInfo.FrozenBy,
			FrozenCause: frozenUserInfo.FrozenCause,
		}).
		Post("http://localhost:50070/v1/frozen/user")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp2.StatusCode())
		info := npool.FrozenUserResponse{}
		err := json.Unmarshal(resp2.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.Info.ID, uuid.UUID{})
			assert.Equal(t, info.Info.UserID, frozenUserInfo.UserID)
			assert.Equal(t, info.Info.FrozenBy, frozenUserInfo.FrozenBy)
			assert.Equal(t, info.Info.FrozenCause, frozenUserInfo.FrozenCause)
			frozenUserInfo.ID = info.Info.ID
		}
	}

	resp5, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.QueryUserFrozenRequest{
			UserID: frozenUserInfo.UserID,
		}).
		Post("http://localhost:50070/v1/query/user/frozen")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp5.StatusCode())
		info := npool.QueryUserFrozenResponse{}
		err := json.Unmarshal(resp5.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.UserID, frozenUserInfo.UserID)
			assert.Equal(t, info.Info.FrozenBy, frozenUserInfo.FrozenBy)
			assert.Equal(t, info.Info.FrozenCause, frozenUserInfo.FrozenCause)
		}
	}

	resp3, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.UnfrozenUserRequest{
			ID:         frozenUserInfo.ID,
			UserID:     frozenUserInfo.UserID,
			UnfrozenBy: frozenUserInfo.FrozenBy,
		}).
		Post("http://localhost:50070/v1/unfrozen/user")
	if assert.Nil(t, err) {
		fmt.Println("resp3 is", resp3)
		assert.Equal(t, 200, resp3.StatusCode())
		info := npool.UnfrozenUserResponse{}
		err := json.Unmarshal(resp3.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, frozenUserInfo.ID)
			assert.Equal(t, info.Info.UserID, frozenUserInfo.UserID)
			assert.Equal(t, info.Info.FrozenBy, frozenUserInfo.FrozenBy)
			assert.Equal(t, info.Info.FrozenCause, frozenUserInfo.FrozenCause)
			assert.Equal(t, info.Info.UnfrozenBy, frozenUserInfo.FrozenBy)
		}
	}

	resp4, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.GetFrozenUsersRequest{}).
		Post("http://localhost:50070/v1/get/frozen/user")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp4.StatusCode())
	}
}
