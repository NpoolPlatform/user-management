package userprovider

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/user"
	userinfo "github.com/NpoolPlatform/user-management/pkg/crud/user-info"
	userprovider "github.com/NpoolPlatform/user-management/pkg/crud/user-provider"
)

func BindThirdParty(ctx context.Context, in *npool.BindThirdPartyRequest) (*npool.BindThirdPartyResponse, error) {
	_, err := userinfo.QueryUserByUserID(ctx, in.UserID)
	if err != nil {
		return nil, err
	}

	resp, err := userprovider.Create(ctx, in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetUserProviders(ctx context.Context, in *npool.GetUserProvidersRequest) (*npool.GetUserProvidersResponse, error) {
	_, err := userinfo.QueryUserByUserID(ctx, in.UserID)
	if err != nil {
		return nil, err
	}

	resp, err := userprovider.Get(ctx, in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func UnbindUserProviders(ctx context.Context, in *npool.UnbindThirdPartyRequest) (*npool.UnbindThirdPartyResponse, error) {
	_, err := userinfo.QueryUserByUserID(ctx, in.UserID)
	if err != nil {
		return nil, err
	}
	resp, err := userprovider.Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func QueryUserByUserProviderID(ctx context.Context, in *npool.QueryUserByUserProviderIDRequest) (*npool.QueryUserByUserProviderIDResponse, error) {
	resp, err := userprovider.QueryUserProviderInfoByProviderUserID(ctx, in)
	if err != nil {
		return nil, err
	}

	userBasicInfo, err := userinfo.QueryUserByUserID(ctx, resp.Info.UserProviderInfo.UserID)
	if err != nil {
		return nil, err
	}
	resp.Info.UserBasicInfo = userBasicInfo
	return resp, nil
}
