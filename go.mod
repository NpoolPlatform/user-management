module github.com/NpoolPlatform/user-management

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/NpoolPlatform/application-management v0.0.0-20211117021047-b393a59181fa
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211117074545-bc1340849b08
	github.com/NpoolPlatform/verification-door v0.0.0-20211117021007-8fcff5172d94
	github.com/aokoli/goutils v1.1.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.6.2 // indirect
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mwitkow/go-proto-validators v0.3.2 // indirect
	github.com/pseudomuto/protoc-gen-doc v1.5.0 // indirect
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/crypto v0.0.0-20211117183948-ae814b36b871
	golang.org/x/sys v0.0.0-20211116061358-0a5406a5449c // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/genproto v0.0.0-20211117155847-120650a500bb
	google.golang.org/grpc v1.41.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
