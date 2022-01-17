package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/user"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedUserServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterUserServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterUserHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
