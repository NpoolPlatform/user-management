package userinfo

import (
	"context"
	"time"

	"github.com/NpoolPlatform/user-management/message/npool"
	"github.com/NpoolPlatform/user-management/pkg/db"
	"github.com/NpoolPlatform/user-management/pkg/db/ent"
	"github.com/NpoolPlatform/user-management/pkg/db/ent/user"
	"github.com/NpoolPlatform/user-management/pkg/encryption"
	"github.com/NpoolPlatform/user-management/pkg/utils"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

const UserExistError = "user doesn't exist"

func dbRowToInfo(row *ent.User) *npool.UserBasicInfo {
	return &npool.UserBasicInfo{
		UserID:         row.ID.String(),
		Username:       row.Username,
		Password:       row.Password,
		DisplayName:    row.DisplayName,
		PhoneNumber:    row.PhoneNumber,
		EmailAddress:   row.EmailAddress,
		SignupMethod:   row.SignupMethod,
		CreateAt:       row.CreateAt,
		UpdateAt:       row.UpdateAt,
		Avatar:         row.Avatar,
		Region:         row.Region,
		Age:            row.Age,
		Gender:         row.Gender,
		Birthday:       row.Birthday,
		Country:        row.Country,
		Province:       row.Province,
		City:           row.City,
		Career:         row.Career,
		FirstName:      row.FirstName,
		LastName:       row.LastName,
		StreetAddress1: row.StreetAddress1,
		StreetAddress2: row.StreetAddress2,
		Compony:        row.Compony,
	}
}

func Create(ctx context.Context, in *npool.AddUserRequest) (*npool.AddUserResponse, error) {
	password := in.GetUserInfo().GetPassword()

	salt := encryption.Salt()
	hashPassword, err := encryption.EncryptePassword(password, salt)
	if err != nil {
		return nil, err
	}

	info, err := db.Client().
		User.
		Create().
		SetUsername(in.UserInfo.Username).
		SetPassword(hashPassword).
		SetSalt(salt).
		SetDisplayName(in.UserInfo.DisplayName).
		SetPhoneNumber(in.UserInfo.PhoneNumber).
		SetEmailAddress(in.UserInfo.EmailAddress).
		SetSignupMethod(in.UserInfo.SignupMethod).
		SetAvatar(in.UserInfo.Avatar).
		SetRegion(in.UserInfo.Region).
		SetAge(in.UserInfo.Age).
		SetGender(in.UserInfo.Gender).
		SetBirthday(in.UserInfo.Birthday).
		SetCountry(in.UserInfo.Country).
		SetProvince(in.UserInfo.Province).
		SetCity(in.UserInfo.City).
		SetCareer(in.UserInfo.Career).
		SetFirstName(in.UserInfo.FirstName).
		SetLastName(in.UserInfo.LastName).
		SetStreetAddress1(in.UserInfo.StreetAddress1).
		SetStreetAddress2(in.UserInfo.StreetAddress2).
		SetCompony(in.UserInfo.Compony).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to create user: %v", err)
	}

	return &npool.AddUserResponse{
		Info: dbRowToInfo(info),
	}, nil
}

func SetPassword(ctx context.Context, password, userID string) error {
	id, err := uuid.Parse(userID)
	if err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}

	salt := encryption.Salt()
	hashPassword, err := encryption.EncryptePassword(password, salt)
	if err != nil {
		return err
	}

	_, err = db.Client().
		User.
		UpdateOneID(id).
		SetPassword(hashPassword).
		SetSalt(salt).
		Save(ctx)
	if err != nil {
		return xerrors.Errorf("fail to set password: %v", err)
	}
	return nil
}

func Update(ctx context.Context, in *npool.UpdateUserInfoRequest) (*npool.UpdateUserInfoResponse, error) {
	id, err := uuid.Parse(in.Info.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	ok := utils.RegexpUsername(in.Info.Username)
	if !ok {
		return nil, xerrors.Errorf("username not legal")
	}

	resp, err := QueryUserByParam(ctx, in.Info.Username)
	if err == nil && resp.UserID != in.Info.UserID {
		return nil, xerrors.Errorf("username has been used")
	}

	info, err := db.Client().User.
		UpdateOneID(id).
		SetUsername(in.Info.Username).
		SetAvatar(in.Info.Avatar).
		SetRegion(in.Info.Region).
		SetAge(in.Info.Age).
		SetGender(in.Info.Gender).
		SetBirthday(in.Info.Birthday).
		SetCountry(in.Info.Country).
		SetProvince(in.Info.Province).
		SetCity(in.Info.City).
		SetCareer(in.Info.Career).
		SetDisplayName(in.Info.DisplayName).
		SetFirstName(in.Info.FirstName).
		SetLastName(in.Info.LastName).
		SetStreetAddress1(in.Info.StreetAddress1).
		SetStreetAddress2(in.Info.StreetAddress2).
		SetCompony(in.Info.Compony).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update user info: %v", err)
	}

	return &npool.UpdateUserInfoResponse{
		Info: dbRowToInfo(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetUserRequest) (*npool.GetUserResponse, error) {
	id, err := uuid.Parse(in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	info, err := db.Client().
		User.
		Query().
		Where(
			user.And(
				user.ID(id),
				user.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query user info: %v", err)
	}

	if len(info) == 0 {
		return nil, xerrors.Errorf(UserExistError)
	}

	return &npool.GetUserResponse{
		Info: dbRowToInfo(info[0]),
	}, nil
}

func GetAll(ctx context.Context) (*npool.GetUsersResponse, error) {
	infos, err := db.Client().
		User.
		Query().
		Where(
			user.And(
				user.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query users: %v", err)
	}
	response := []*npool.UserBasicInfo{}
	for _, info := range infos {
		response = append(response, dbRowToInfo(info))
	}
	return &npool.GetUsersResponse{
		Infos: response,
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteUserRequest) (*npool.DeleteUserResponse, error) {
	for _, id := range in.DeleteUserIDs {
		id, err := uuid.Parse(id)
		if err != nil {
			return nil, xerrors.Errorf("invalid user id: %v", err)
		}

		_, err = db.Client().
			User.
			UpdateOneID(id).
			SetDeleteAt(uint32(time.Now().Unix())).
			Save(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail to delete user: %v", err)
		}
	}
	return &npool.DeleteUserResponse{
		Info: "delete users successfully",
	}, nil
}

func GetUserPassword(ctx context.Context, userID string) (string, error) {
	id, err := uuid.Parse(userID)
	if err != nil {
		return "", xerrors.Errorf("invalid user id: %v", err)
	}
	info, err := db.Client().
		User.
		Query().
		Where(
			user.And(
				user.ID(id),
				user.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return "", xerrors.Errorf("fail to query user password: %v", err)
	}

	if len(info) == 0 {
		return "", xerrors.Errorf(UserExistError)
	}
	return info[0].Password, nil
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
			user.And(
				user.ID(id),
				user.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return "", xerrors.Errorf("fail to query user info: %v", err)
	}

	if len(info) == 0 {
		return "", xerrors.Errorf(UserExistError)
	}

	return info[0].Salt, nil
}

func QueryUserByParam(ctx context.Context, param string) (*npool.UserBasicInfo, error) {
	info, err := db.Client().
		User.
		Query().
		Where(
			user.DeleteAt(0),
			user.Or(
				user.Username(param),
				user.EmailAddress(param),
				user.PhoneNumber(param),
			),
		).Only(ctx)
	if err != nil {
		return &npool.UserBasicInfo{}, xerrors.Errorf("fail to get user info: %v", err)
	}

	return dbRowToInfo(info), nil
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
			user.And(
				user.ID(id),
				user.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return &npool.UserBasicInfo{}, xerrors.Errorf("fail to get user info by username: %v", err)
	}

	if len(info) == 0 {
		return nil, xerrors.Errorf(UserExistError)
	}

	return dbRowToInfo(info[0]), nil
}

func QueryUserExist(ctx context.Context, in *npool.QueryUserExistRequest) (*npool.QueryUserExistResponse, error) {
	info, err := db.Client().
		User.
		Query().
		Where(
			user.DeleteAt(0),
			user.Or(
				user.Username(in.Username),
				user.PhoneNumber(in.Username),
				user.EmailAddress(in.Username),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query user: %v", err)
	}

	if in.Password != "" {
		salt, err := GetUserSalt(ctx, info.ID.String())
		if err != nil {
			return nil, xerrors.Errorf("fail to get user's salt: %v", err)
		}

		dbPassword, err := GetUserPassword(ctx, info.ID.String())
		if err != nil {
			return nil, xerrors.Errorf("fail to get user's password: %v", err)
		}

		err = encryption.VerifyUserPassword(in.Password, dbPassword, salt)
		if err != nil {
			return nil, xerrors.Errorf("user password not equal: %v", err)
		}
	}
	return &npool.QueryUserExistResponse{
		Info: dbRowToInfo(info),
	}, nil
}

func SetUserPhone(ctx context.Context, userID, phone string) error {
	id, err := uuid.Parse(userID)
	if err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	resp, err := QueryUserByParam(ctx, phone)
	if err == nil && resp.UserID != userID {
		return xerrors.Errorf("phone number has been used")
	}

	_, err = db.Client().
		User.
		UpdateOneID(id).
		SetPhoneNumber(phone).
		Save(ctx)
	if err != nil {
		return xerrors.Errorf("fail to set phone number into database: %v", err)
	}

	return nil
}

func SetUserEmail(ctx context.Context, userID, email string) error {
	id, err := uuid.Parse(userID)
	if err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	resp, err := QueryUserByParam(ctx, email)
	if err == nil && resp.UserID != userID {
		return xerrors.Errorf("email address has been used")
	}

	_, err = db.Client().
		User.
		UpdateOneID(id).
		SetEmailAddress(email).
		Save(ctx)
	if err != nil {
		return xerrors.Errorf("fail to set email address into database: %v", err)
	}

	return nil
}
