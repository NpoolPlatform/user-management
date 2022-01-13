package grpc

import (
	"context"
	"time"

	applicationconst "github.com/NpoolPlatform/application-management/pkg/message/const"
	mygrpc "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	pbApplication "github.com/NpoolPlatform/message/npool/application"
	pbVerification "github.com/NpoolPlatform/message/npool/verification"
	verificationconst "github.com/NpoolPlatform/verification-door/pkg/message/const"
	"golang.org/x/xerrors"
)

const (
	grpcTimeout = 5 * time.Second
)

func VerifyCode(ctx context.Context, param, code string) error {
	conn, err := mygrpc.GetGRPCConn(verificationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbVerification.NewVerificationDoorClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	_, err = client.VerifyCode(ctx, &pbVerification.VerifyCodeRequest{
		Param: param,
		Code:  code,
	})
	if err != nil {
		return err
	}
	return nil
}

func VerifyCodeWithUserID(ctx context.Context, userID, param, code string) error {
	conn, err := mygrpc.GetGRPCConn(verificationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbVerification.NewVerificationDoorClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	_, err = client.VerifyCodeWithUserID(ctx, &pbVerification.VerifyCodeWithUserIDRequest{
		UserID: userID,
		Code:   code,
		Param:  param,
	})
	if err != nil {
		return err
	}
	return nil
}

func VerifyGoogleCode(ctx context.Context, userID, appID, code string) error {
	conn, err := mygrpc.GetGRPCConn(verificationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbVerification.NewVerificationDoorClient(conn)
	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	_, err = client.VerifyGoogleAuth(ctx, &pbVerification.VerifyGoogleAuthRequest{
		AppID:  appID,
		UserID: userID,
		Code:   code,
	})
	if err != nil {
		return xerrors.Errorf("fail to verify google authentication: %v", err)
	}
	return nil
}

func QueryAppExist(ctx context.Context, appID string) error {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbApplication.NewApplicationManagementClient(conn)
	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	resp, err := client.GetApplication(ctx, &pbApplication.GetApplicationRequest{
		AppID: appID,
	})
	if err != nil {
		return err
	}
	if resp != nil {
		return nil
	}
	return xerrors.Errorf("app not exist")
}

func AddUserToApplication(ctx context.Context, userID, appID string) error {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbApplication.NewApplicationManagementClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	_, err = client.AddUsersToApplication(ctx, &pbApplication.AddUsersToApplicationRequest{
		UserIDs: []string{userID},
		AppID:   appID,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetUserApplicationInfo(ctx context.Context, userID, appID string) (*pbApplication.ApplicationUserDetail, error) {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbApplication.NewApplicationManagementClient(conn)
	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	resp, err := client.GetApplicationUserDetail(ctx, &pbApplication.GetApplicationUserDetailRequest{
		UserID: userID,
		AppID:  appID,
	})
	if err != nil {
		return nil, err
	}
	return resp.Info, nil
}
