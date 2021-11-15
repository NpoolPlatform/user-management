package grpc

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/user-management/pkg/test-init"
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

func TestGrpc(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	client, err := newVerificationGrpcClient()
	if assert.Nil(t, err) {
		assert.NotNil(t, client)
	}

	email := "crazyzplzpl@163.com"

	err = VerifyCode(email, "12345")
	assert.NotNil(t, err)
}
