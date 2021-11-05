package userinfo

import (
	"context"
	"time"

	"github.com/NpoolPlatform/user-management/message/npool"
	"github.com/NpoolPlatform/user-management/pkg/db"
	"github.com/NpoolPlatform/user-management/pkg/db/ent"
	"github.com/NpoolPlatform/user-management/pkg/db/ent/user"
	"github.com/NpoolPlatform/user-management/pkg/encryption"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func dbRowToInfo(row *ent.User) *npool.UserBasicInfo {
	return &npool.UserBasicInfo{
		UserId:       row.ID.String(),
		Username:     row.Username,
		Password:     row.Password,
		DisplayName:  row.DisplayName,
		PhoneNumber:  row.PhoneNumber,
		EmailAddress: row.EmailAddress,
		LoginTimes:   row.LoginTimes,
		KycVerify:    row.KycVerify,
		GaVerify:     row.GaVerify,
		SignupMethod: row.SignupMethod,
		CreateAt:     int32(row.CreateAt),
		UpdateAt:     int32(row.UpdateAt),
		Avatar:       row.Avatar,
		Region:       row.Region,
		Age:          row.Age,
		Gender:       row.Gender,
		Birthday:     row.Birthday,
		Country:      row.Country,
		Province:     row.Province,
		City:         row.City,
		Career:       row.Career,
	}
}

func Create(ctx context.Context, in *npool.AddUserRequest) (*npool.AddUserResponse, error) {
	myRow, err := db.Client().User.Query().Where(
		user.Or(
			user.Username(in.GetUserInfo().GetUsername()),
		),
	).All(ctx)
	if err != nil {
		return nil, err
	}

	if len(myRow) != 0 {
		if myRow[0].DeleteAt != 0 {
			return nil, xerrors.Errorf("this user has already existed and has been deleted!")
		}
		return nil, xerrors.Errorf("user has already exist!")
	}

	password := in.GetUserInfo().GetPassword()

	salt := encryption.Salt()
	hashPassword, err := encryption.EncryptePassword(password, salt)
	if err != nil {
		return nil, err
	}
	info, err := db.Client().User.Create().
		SetID(uuid.MustParse(in.UserInfo.GetUserId())).
		SetUsername(in.GetUserInfo().GetUsername()).
		SetPassword(hashPassword).
		SetSalt(salt).
		SetDisplayName(in.GetUserInfo().GetDisplayName()).
		SetPhoneNumber(in.GetUserInfo().GetPhoneNumber()).
		SetEmailAddress(in.GetUserInfo().GetEmailAddress()).
		SetLoginTimes(in.GetUserInfo().GetLoginTimes()).
		SetKycVerify(in.GetUserInfo().GetKycVerify()).
		SetGaVerify(in.GetUserInfo().GetGaVerify()).
		SetSignupMethod(in.GetUserInfo().GetSignupMethod()).
		SetAvatar(in.GetUserInfo().GetAvatar()).
		SetRegion(in.GetUserInfo().GetRegion()).
		SetAge(in.GetUserInfo().GetAge()).
		SetGender(in.GetUserInfo().Gender).
		SetBirthday(in.GetUserInfo().GetBirthday()).
		SetCountry(in.GetUserInfo().GetCountry()).
		SetProvince(in.GetUserInfo().GetProvince()).
		SetCity(in.GetUserInfo().GetCity()).
		SetCareer(in.GetUserInfo().GetCareer()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to create user: %v", err)
	}

	return &npool.AddUserResponse{
		UserInfo: dbRowToInfo(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateUserInfoRequest) (*npool.UpdateUserInfoResponse, error) {
	id, err := uuid.Parse(in.GetUserInfo().GetUserId())
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}
	myRow, err := db.Client().User.Query().Where(
		user.Or(
			user.ID(id),
		),
	).All(ctx)
	if err != nil {
		return nil, err
	}

	if len(myRow) == 0 {
		return nil, xerrors.Errorf("user doesn't exist")
	}

	if myRow[0].DeleteAt != 0 {
		return nil, xerrors.Errorf("user has already been deleted!")
	}

	info, err := db.Client().User.
		UpdateOneID(id).
		SetAvatar(in.GetUserInfo().GetAvatar()).
		SetRegion(in.GetUserInfo().GetRegion()).
		SetAge(in.GetUserInfo().GetAge()).
		SetGender(in.GetUserInfo().Gender).
		SetBirthday(in.GetUserInfo().GetBirthday()).
		SetCountry(in.GetUserInfo().GetCountry()).
		SetProvince(in.GetUserInfo().GetProvince()).
		SetCity(in.GetUserInfo().GetCity()).
		SetCareer(in.GetUserInfo().GetCareer()).
		SetPhoneNumber(in.GetUserInfo().GetPhoneNumber()).
		SetEmailAddress(in.GetUserInfo().GetEmailAddress()).
		SetLoginTimes(in.GetUserInfo().GetLoginTimes()).
		SetKycVerify(in.GetUserInfo().GetKycVerify()).
		SetGaVerify(in.GetUserInfo().GetGaVerify()).
		SetDisplayName(in.GetUserInfo().GetDisplayName()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update user info: %v", err)
	}

	return &npool.UpdateUserInfoResponse{
		UserInfo: dbRowToInfo(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetUserRequest) (*npool.GetUserResponse, error) {
	id, err := uuid.Parse(in.UserId)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	info, err := db.Client().
		User.
		Query().
		Where(
			user.Or(
				user.ID(id),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query user info: %v", err)
	}

	if len(info) == 0 {
		return nil, xerrors.Errorf("empty result of user info")
	}

	if info[0].DeleteAt != 0 {
		return nil, xerrors.Errorf("user has been deleted")
	}

	return &npool.GetUserResponse{
		UserInfo: dbRowToInfo(info[0]),
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteUserRequest) (*npool.DeleteUserResponse, error) {
	for _, id := range in.DeleteUserIds {
		id, err := uuid.Parse(id)
		if err != nil {
			return nil, xerrors.Errorf("invalid user id: %v", err)
		}

		_, err = db.Client().
			User.
			UpdateOneID(id).
			SetDeleteAt(time.Now().UnixNano()).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail to delete user: %v", err)
		}
	}
	return &npool.DeleteUserResponse{
		Info: "delete users successfully",
	}, nil
}

func GetUserSalt(ctx context.Context, userID string) (string, error) {
	id, err := uuid.Parse(userID)
	if err != nil {
		return "", xerrors.Errorf("invalid user id: %v", err)
	}
	info, err := db.Client().
		User.
		Query().
		Where(
			user.Or(
				user.ID(id),
			),
		).All(ctx)
	if err != nil {
		return "", xerrors.Errorf("fail to query user info: %v", err)
	}

	if len(info) == 0 {
		return "", xerrors.Errorf("user is not exist")
	}

	if info[0].DeleteAt != 0 {
		return "", xerrors.Errorf("user has already deleted")
	}

	return info[0].Salt, nil
}

func QueryUserByUsername(ctx context.Context, username string) (*npool.UserBasicInfo, error) {
	info, err := db.Client().
		User.
		Query().
		Where(
			user.Or(
				user.Username(username),
				user.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return &npool.UserBasicInfo{}, xerrors.Errorf("fail to get user info by username: %v", err)
	}

	if len(info) == 0 {
		return &npool.UserBasicInfo{}, xerrors.Errorf("user doesn't exist")
	}
	return dbRowToInfo(info[0]), nil
}

func QueryUserByUserID(ctx context.Context, userID string) (*npool.UserBasicInfo, error) {
	id, err := uuid.Parse(userID)
	if err != nil {
		return &npool.UserBasicInfo{}, xerrors.Errorf("fail to invalid user id: %v", err)
	}

	info, err := db.Client().
		User.
		Query().
		Where(
			user.Or(
				user.ID(id),
				user.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return &npool.UserBasicInfo{}, xerrors.Errorf("fail to get user info by username: %v", err)
	}

	if len(info) == 0 {
		return &npool.UserBasicInfo{}, xerrors.Errorf("user doesn't exist")
	}
	return dbRowToInfo(info[0]), nil
}

func QueryUserByPhoneNumber(ctx context.Context, phoneNumber string) (*npool.UserBasicInfo, error) {
	info, err := db.Client().
		User.
		Query().
		Where(
			user.Or(
				user.PhoneNumber(phoneNumber),
				user.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return &npool.UserBasicInfo{}, xerrors.Errorf("fail to get user info by username: %v", err)
	}

	if len(info) == 0 {
		return &npool.UserBasicInfo{}, xerrors.Errorf("user doesn't exist")
	}
	return dbRowToInfo(info[0]), nil
}

func QueryUserByEmailAddress(ctx context.Context, emailAddress string) (*npool.UserBasicInfo, error) {
	info, err := db.Client().
		User.
		Query().
		Where(
			user.Or(
				user.EmailAddress(emailAddress),
				user.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return &npool.UserBasicInfo{}, xerrors.Errorf("fail to get user info by username: %v", err)
	}

	if len(info) == 0 {
		return &npool.UserBasicInfo{}, xerrors.Errorf("user doesn't exist")
	}
	return dbRowToInfo(info[0]), nil
}
