package userprovider

import (
	"context"
	"time"

	"github.com/NpoolPlatform/user-management/message/npool"
	"github.com/NpoolPlatform/user-management/pkg/db"
	"github.com/NpoolPlatform/user-management/pkg/db/ent"
	"github.com/NpoolPlatform/user-management/pkg/db/ent/userprovider"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type ProviderUserInfo struct {
	OauthAvatarURL   string `json:"OauthAvatarURL"`
	OauthDisplayName string `json:"OauthDisplayName"`
	OauthEmail       string `json:"OauthEmail"`
	OauthID          string `json:"OauthID"`
	OauthUsername    string `json:"OauthUsername"`
}

func dbRowToInfo(row *ent.UserProvider) *npool.UserProvider {
	return &npool.UserProvider{
		ID:               row.ID.String(),
		UserId:           row.UserID.String(),
		ProviderId:       row.ProviderID.String(),
		ProviderUserId:   row.ProviderUserID,
		UserProviderInfo: row.UserProviderInfo,
		CreateAt:         row.CreateAt,
		UpdateAt:         row.UpdateAt,
	}
}

func Create(ctx context.Context, in *npool.BindThirdPartyRequest) (*npool.BindThirdPartyResponse, error) {
	userID, err := uuid.Parse(in.UserId)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}
	providerID, err := uuid.Parse(in.ProviderId)
	if err != nil {
		return nil, xerrors.Errorf("invalid provider id: %v", err)
	}
	info, err := db.Client().
		UserProvider.
		Query().
		Where(
			userprovider.And(
				userprovider.UserID(userID),
				userprovider.ProviderID(providerID),
				userprovider.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get user provider info: %v", err)
	}

	if len(info) != 0 {
		return nil, xerrors.Errorf("user has been binded this provider")
	}

	createInfo, err := db.Client().
		UserProvider.
		Create().
		SetUserID(userID).
		SetProviderID(providerID).
		SetProviderUserID(in.ProviderUserId).
		SetUserProviderInfo(in.UserProviderInfo).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to bind third party provider: %v", err)
	}
	return &npool.BindThirdPartyResponse{
		Info: dbRowToInfo(createInfo),
	}, nil
}

func Get(ctx context.Context, in *npool.GetUserProvidersRequest) (*npool.GetUserProvidersResponse, error) {
	userID, err := uuid.Parse(in.UserId)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}
	infos, err := db.Client().
		UserProvider.
		Query().
		Where(
			userprovider.And(
				userprovider.UserID(userID),
				userprovider.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query user provider list: %v", err)
	}

	userProviders := []*npool.UserProvider{}
	for _, info := range infos {
		userProviders = append(userProviders, dbRowToInfo(info))
	}
	return &npool.GetUserProvidersResponse{
		Infos: userProviders,
	}, nil
}

func Delete(ctx context.Context, in *npool.UnbindThirdPartyRequest) (*npool.UnbindThirdPartyResponse, error) {
	userID, err := uuid.Parse(in.UserId)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	providerID, err := uuid.Parse(in.ProviderId)
	if err != nil {
		return nil, xerrors.Errorf("invalid provider id: %v", err)
	}

	id, providerUserID, err := QueryUserProviderIDByUserIDAndProviderID(ctx, userID, providerID)
	if err != nil {
		return nil, err
	}

	info, err := db.Client().
		UserProvider.
		UpdateOneID(id).
		SetProviderUserID("deleted" + providerUserID + time.Now().String()).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to unbind user provider: %v", err)
	}

	return &npool.UnbindThirdPartyResponse{
		Info: dbRowToInfo(info),
	}, nil
}

func QueryUserProviderIDByUserIDAndProviderID(ctx context.Context, userID, providerID uuid.UUID) (uuid.UUID, string, error) {
	info, err := db.Client().
		UserProvider.
		Query().
		Where(
			userprovider.And(
				userprovider.UserID(userID),
				userprovider.ProviderID(providerID),
			),
		).All(ctx)
	if err != nil {
		return uuid.UUID{}, "", xerrors.Errorf("fail to get user provider info: %v", err)
	}
	if len(info) == 0 {
		return uuid.UUID{}, "", xerrors.Errorf("empty user provider")
	}
	return info[0].ID, info[0].ProviderUserID, nil
}

func QueryUserProviderInfoByProviderUserID(ctx context.Context, in *npool.QueryUserByUserProviderIDRequest) (*npool.QueryUserByUserProviderIDResponse, error) {
	providerID, err := uuid.Parse(in.ProviderID)
	if err != nil {
		return nil, xerrors.Errorf("invalid provider id: %v", err)
	}
	info, err := db.Client().
		UserProvider.
		Query().
		Where(
			userprovider.And(
				userprovider.ProviderUserID(in.ProviderUserID),
				userprovider.ProviderID(providerID),
				userprovider.DeleteAt(0),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get user provider info by user provider id: %v", err)
	}
	return &npool.QueryUserByUserProviderIDResponse{
		Info: &npool.QueryProviderUserInfo{
			UserProviderInfo: dbRowToInfo(info),
		},
	}, nil
}
