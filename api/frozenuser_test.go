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

	createUser := npool.UserBasicInfo{
		Username:     "test-frozen" + uuid.New().String(),
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
		Post("http://localhost:32759/v1/signup")
	fmt.Println("sign up error", err)
	if assert.Nil(t, err) {
		fmt.Println("resp1 is", resp1)
		assert.Equal(t, 200, resp1.StatusCode())
		info := npool.SignupResponse{}
		err := json.Unmarshal(resp1.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.UserInfo.UserId, uuid.UUID{})
			assertUserInfo(t, info.UserInfo, &createUser)
			createUser.UserId = info.UserInfo.UserId
		}
	}

	frozenUserInfo := npool.FrozenUser{
		UserId:      createUser.UserId,
		FrozenBy:    uuid.New().String(),
		FrozenCause: "user has done some illegal operations",
	}

	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.FrozenUserRequest{
			UserId:      frozenUserInfo.UserId,
			FrozenBy:    frozenUserInfo.FrozenBy,
			FrozenCause: frozenUserInfo.FrozenCause,
		}).
		Post("http://localhost:32759/v1/frozen/user")
	if assert.Nil(t, err) {
		fmt.Println("resp2 is", resp2)
		assert.Equal(t, 200, resp2.StatusCode())
		info := npool.FrozenUserResponse{}
		err := json.Unmarshal(resp2.Body(), &info)
		if assert.Nil(t, err) {
			assert.NotEqual(t, info.FrozenUserInfo.Id, uuid.UUID{})
			assert.Equal(t, info.FrozenUserInfo.UserId, frozenUserInfo.UserId)
			assert.Equal(t, info.FrozenUserInfo.FrozenBy, frozenUserInfo.FrozenBy)
			assert.Equal(t, info.FrozenUserInfo.FrozenCause, frozenUserInfo.FrozenCause)
			frozenUserInfo.Id = info.FrozenUserInfo.Id
		}
	}

	resp3, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.UnfrozenUserRequest{
			Id:         frozenUserInfo.Id,
			UserId:     frozenUserInfo.UserId,
			UnfrozenBy: frozenUserInfo.FrozenBy,
		}).
		Post("http://localhost:32759/v1/unfrozen/user")
	if assert.Nil(t, err) {
		fmt.Println("resp3 is", resp3)
		assert.Equal(t, 200, resp3.StatusCode())
		info := npool.UnfrozenUserResponse{}
		err := json.Unmarshal(resp3.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.UnFrozenUserInfo.Id, frozenUserInfo.Id)
			assert.Equal(t, info.UnFrozenUserInfo.UserId, frozenUserInfo.UserId)
			assert.Equal(t, info.UnFrozenUserInfo.FrozenBy, frozenUserInfo.FrozenBy)
			assert.Equal(t, info.UnFrozenUserInfo.FrozenCause, frozenUserInfo.FrozenCause)
			assert.Equal(t, info.UnFrozenUserInfo.UnfrozenBy, frozenUserInfo.FrozenBy)
		}
	}

	resp4, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.GetFrozenUsersRequest{}).
		Post("http://localhost:32759/v1/get/frozen/user")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp4.StatusCode())
	}
}
