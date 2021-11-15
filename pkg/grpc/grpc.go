package grpc

import (
	"context"
	"strings"

	pbApplication "github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	mygrpc "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	pbVerification "github.com/NpoolPlatform/verification-door/message/npool"
	"google.golang.org/grpc"
)

const (
	VerificationService = "verification-door.npool.top"
	ApplicationService  = "application-management.npool.top"
)

func newVerificationGrpcClient() (*grpc.ClientConn, error) {
	serviceAgent, err := config.PeekService(VerificationService)
	if err != nil {
		return nil, err
	}

	myAddress := []string{}
	for _, address := range strings.Split(serviceAgent.Address, ",") {
		myAddress = append(myAddress, address+":50091")
	}

	conn, err := mygrpc.GetGRPCConn(strings.Join(myAddress, ","))
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
	serviceAgent, err := config.PeekService(ApplicationService)
	if err != nil {
		return nil, err
	}

	myAddress := []string{}
	for _, address := range strings.Split(serviceAgent.Address, ",") {
		myAddress = append(myAddress, address+":50081")
	}

	conn, err := mygrpc.GetGRPCConn(strings.Join(myAddress, ","))
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
