package grpc

import (
	"context"
	"fmt"
	"net"

	pbApplication "github.com/NpoolPlatform/application-management/message/npool"
	applicationconst "github.com/NpoolPlatform/application-management/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	mygrpc "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	pbVerification "github.com/NpoolPlatform/verification-door/message/npool"
	verificationconst "github.com/NpoolPlatform/verification-door/pkg/message/const"
	"google.golang.org/grpc"
)

const (
	VerificationService     = verificationconst.ServiceName
	VerificationServicePort = ":50091"
	ApplicationService      = applicationconst.ServiceName
	ApplicationServicePort  = ":50081"
)

func newVerificationGrpcClient() (*grpc.ClientConn, error) {
	serviceAgent, err := config.PeekService(verificationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	conn, err := mygrpc.GetGRPCConn(net.JoinHostPort(serviceAgent.Address, fmt.Sprintf("%d", serviceAgent.Port)))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func VerifyCode(param, code string) error {
	conn, err := newVerificationGrpcClient()
	if err != nil {
		return err
	}

	defer conn.Close()

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

func newApplicationGrpcClient() (*grpc.ClientConn, error) {
	serviceAgent, err := config.PeekService(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	conn, err := mygrpc.GetGRPCConn(net.JoinHostPort(serviceAgent.Address, fmt.Sprintf("%d", serviceAgent.Port)))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func AddUserToApplication(userID, appID string) error {
	conn, err := newApplicationGrpcClient()
	if err != nil {
		return err
	}
	defer conn.Close()

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
