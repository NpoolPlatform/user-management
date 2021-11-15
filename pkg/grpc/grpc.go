package grpc

import (
	"context"
	"strings"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	mygrpc "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	pb "github.com/NpoolPlatform/verification-door/message/npool"
	"google.golang.org/grpc"
)

const (
	VerificationService = "verification-door.npool.top"
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

	client := pb.NewVerificationDoorClient(conn)
	_, err = client.VerifyCode(context.Background(), &pb.VerifyCodeRequest{
		Param: param,
		Code:  code,
	})
	if err != nil {
		return err
	}
	return err
}
