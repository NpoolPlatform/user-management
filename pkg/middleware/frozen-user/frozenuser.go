package frozenuser

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/user"
	frozeninfo "github.com/NpoolPlatform/user-management/pkg/crud/frozen-info"
	userinfo "github.com/NpoolPlatform/user-management/pkg/crud/user-info"
)

func FrozenUser(ctx context.Context, in *npool.FrozenUserRequest) (*npool.FrozenUserResponse, error) {
	_, err := userinfo.QueryUserByUserID(ctx, in.UserID)
	if err != nil {
		return nil, err
	}

	resp, err := frozeninfo.Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func UnfrozenUser(ctx context.Context, in *npool.UnfrozenUserRequest) (*npool.UnfrozenUserResponse, error) {
	_, err := userinfo.QueryUserByUserID(ctx, in.UserID)
	if err != nil {
		return nil, err
	}
	resp, err := frozeninfo.Update(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
