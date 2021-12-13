module github.com/NpoolPlatform/user-management

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/AmirSoleimani/VoucherCodeGenerator v0.0.0-20201014193813-0206853dccb9
	github.com/NpoolPlatform/application-management v0.0.0-20211209130843-8d9879671402
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211207121121-adb2402676f0
	github.com/NpoolPlatform/verification-door v0.0.0-20211213083447-7a919ad18b36
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.1
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/crypto v0.0.0-20211202192323-5770296d904e
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa
	google.golang.org/grpc v1.42.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
