package frozeninfo

import (
	"context"
	"time"

	"github.com/NpoolPlatform/user-management/message/npool"
	"github.com/NpoolPlatform/user-management/pkg/db"
	"github.com/NpoolPlatform/user-management/pkg/db/ent"
	"github.com/NpoolPlatform/user-management/pkg/db/ent/userfrozen"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

const (
	FrozenStatus   = "frozen"
	UnfrozenStatus = "unfrozen"
)

func dbRowToInfo(row *ent.UserFrozen) *npool.FrozenUser {
	return &npool.FrozenUser{
		ID:          row.ID.String(),
		UserID:      row.UserID.String(),
		FrozenBy:    row.FrozenBy.String(),
		FrozenCause: row.FrozenCause,
		StartAt:     row.CreateAt,
		EndAt:       row.EndAt,
		Status:      row.Status,
		UnfrozenBy:  row.UnfrozenBy.String(),
	}
}

func Create(ctx context.Context, in *npool.FrozenUserRequest) (*npool.FrozenUserResponse, error) {
	id, err := uuid.Parse(in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	frozenBy, err := uuid.Parse(in.FrozenBy)
	if err != nil {
		return nil, xerrors.Errorf("invalid frozen admin id: %v", err)
	}

	infos, err := db.Client().
		UserFrozen.
		Query().
		Where(
			userfrozen.Or(
				userfrozen.UserID(id),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query user: %v", err)
	}

	if len(infos) != 0 {
		for _, info := range infos {
			if info.EndAt == 0 {
				return nil, xerrors.Errorf("user has already been frozened")
			}
		}
	}

	info, err := db.Client().
		UserFrozen.
		Create().
		SetUserID(id).
		SetFrozenBy(frozenBy).
		SetFrozenCause(in.FrozenCause).
		SetStatus(FrozenStatus).
		SetUnfrozenBy(uuid.UUID{}).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to frozen user: %v", err)
	}

	return &npool.FrozenUserResponse{
		Info: dbRowToInfo(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UnfrozenUserRequest) (*npool.UnfrozenUserResponse, error) {
	id, err := uuid.Parse(in.ID)
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	unfrozenBy, err := uuid.Parse(in.UnfrozenBy)
	if err != nil {
		return nil, xerrors.Errorf("invalid unfrozen admin id: %v", err)
	}

	infos, err := db.Client().
		UserFrozen.
		Query().
		Where(
			userfrozen.Or(
				userfrozen.ID(id),
				userfrozen.Status(FrozenStatus),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query frozen user info")
	}

	if len(infos) == 0 {
		return nil, xerrors.Errorf("frozen user info doesn't exist")
	}

	info, err := db.Client().
		UserFrozen.UpdateOneID(id).
		SetUnfrozenBy(unfrozenBy).
		SetStatus(UnfrozenStatus).
		SetEndAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update frozen user info: %v", err)
	}

	return &npool.UnfrozenUserResponse{
		Info: dbRowToInfo(info),
	}, nil
}

func Get(ctx context.Context, in *npool.QueryUserFrozenRequest) (*npool.QueryUserFrozenResponse, error) {
	UserID, err := uuid.Parse(in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	info, err := db.Client().
		UserFrozen.
		Query().
		Where(
			userfrozen.And(
				userfrozen.UserID(UserID),
				userfrozen.Status(FrozenStatus),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query user frozen info: %v", err)
	}

	return &npool.QueryUserFrozenResponse{
		Info: dbRowToInfo(info),
	}, nil
}

func GetAll(ctx context.Context) (*npool.GetFrozenUsersResponse, error) {
	infos, err := db.Client().
		UserFrozen.
		Query().All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query frozen user list: %v", err)
	}

	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty frozen user list")
	}
	myInfos := []*npool.FrozenUser{}
	for _, info := range infos {
		myInfos = append(myInfos, dbRowToInfo(info))
	}
	return &npool.GetFrozenUsersResponse{
		Infos: myInfos,
	}, nil
}
