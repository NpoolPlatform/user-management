module github.com/NpoolPlatform/user-management

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/NpoolPlatform/application-management v0.0.0-20211114130519-bcb592a33694 // indirect
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211110113912-c4859934a03c
	github.com/NpoolPlatform/verification-door v0.0.0-20211115014754-0a08eaafd7b3
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/crypto v0.0.0-20211108221036-ceb1ce70b4fa
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/genproto v0.0.0-20211104193956-4c6863e31247
	google.golang.org/grpc v1.41.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
