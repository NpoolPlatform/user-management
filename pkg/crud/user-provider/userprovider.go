package userprovider

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/user"
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
		UserID:           row.UserID.String(),
		ProviderID:       row.ProviderID.String(),
		ProviderUserID:   row.ProviderUserID,
		UserProviderInfo: row.UserProviderInfo,
		CreateAt:         row.CreateAt,
		UpdateAt:         row.UpdateAt,
	}
}

func Create(ctx context.Context, in *npool.BindThirdPartyRequest) (*npool.BindThirdPartyResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	userID, err := uuid.Parse(in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}
	providerID, err := uuid.Parse(in.ProviderID)
	if err != nil {
		return nil, xerrors.Errorf("invalid provider id: %v", err)
	}
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
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

	createInfo, err := cli.
		UserProvider.
		Create().
		SetUserID(userID).
		SetProviderID(providerID).
		SetProviderUserID(in.ProviderUserID).
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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	userID, err := uuid.Parse(in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	userID, err := uuid.Parse(in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	providerID, err := uuid.Parse(in.ProviderID)
	if err != nil {
		return nil, xerrors.Errorf("invalid provider id: %v", err)
	}

	id, providerUserID, err := QueryUserProviderIDByUserIDAndProviderID(ctx, userID, providerID)
	if err != nil {
		return nil, err
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return uuid.UUID{}, "", xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	providerID, err := uuid.Parse(in.ProviderID)
	if err != nil {
		return nil, xerrors.Errorf("invalid provider id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
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
