module github.com/NpoolPlatform/user-management

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/AmirSoleimani/VoucherCodeGenerator v0.0.0-20201014193813-0206853dccb9
	github.com/NpoolPlatform/application-management v0.0.0-20211209130843-8d9879671402
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211207121121-adb2402676f0
	github.com/NpoolPlatform/verification-door v0.0.0-20211209125023-e6a5a9fcec2e
	github.com/boombuler/barcode v1.0.1 // indirect
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.1
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/t-yuki/gocover-cobertura v0.0.0-20180217150009-aaee18c8195c // indirect
	github.com/tebeka/go2xunit v1.4.10 // indirect
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/crypto v0.0.0-20211202192323-5770296d904e
	golang.org/x/image v0.0.0-20211028202545-6944b10bf410 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/tools v0.1.8 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa
	google.golang.org/grpc v1.42.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
