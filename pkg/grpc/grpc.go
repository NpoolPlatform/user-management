package grpc

import (
	"context"

	pbApplication "github.com/NpoolPlatform/application-management/message/npool"
	applicationconst "github.com/NpoolPlatform/application-management/pkg/message/const"
	mygrpc "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	pbVerification "github.com/NpoolPlatform/verification-door/message/npool"
	verificationconst "github.com/NpoolPlatform/verification-door/pkg/message/const"
)

func VerifyCode(param, code string) error {
	conn, err := mygrpc.GetGRPCConn(verificationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	client := pbVerification.NewVerificationDoorClient(conn)
	_, err = client.VerifyCode(context.Background(), &pbVerification.VerifyCodeRequest{
		Param: param,
		Code:  code,
	})
	if err != nil {
		return err
	}
	return nil
}

func AddUserToApplication(userID, appID string) error {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	client := pbApplication.NewApplicationManagementClient(conn)
	_, err = client.AddUsersToApplication(context.Background(), &pbApplication.AddUsersToApplicationRequest{
		UserIDs: []string{userID},
		AppID:   appID,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetUserApplicationInfo(userID, appID string) (*pbApplication.ApplicationUserDetail, error) {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	client := pbApplication.NewApplicationManagementClient(conn)
	resp, err := client.GetApplicationUserDetail(context.Background(), &pbApplication.GetApplicationUserDetailRequest{
		UserID: userID,
		AppID:  appID,
	})
	if err != nil {
		return nil, err
	}
	return resp.Info, nil
}
