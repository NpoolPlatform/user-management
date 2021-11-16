package testinit

import (
	"fmt"
	"path"
	"runtime"

	applicationconst "github.com/NpoolPlatform/application-management/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	mysqlconst "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	rabbitmqconst "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/const"
	redisconst "github.com/NpoolPlatform/go-service-framework/pkg/redis/const"
	"github.com/NpoolPlatform/user-management/pkg/db"
	servicename "github.com/NpoolPlatform/user-management/pkg/service-name" //nolint
	verificationconst "github.com/NpoolPlatform/verification-door/pkg/message/const"

	"golang.org/x/xerrors"
)

func Init() error {
	_, myPath, _, ok := runtime.Caller(0)
	if !ok {
		return xerrors.Errorf("cannot get source file path")
	}

	appName := path.Base(path.Dir(path.Dir(path.Dir(myPath))))
	configPath := fmt.Sprintf("%s/../../cmd/%v", path.Dir(myPath), appName)

	err := app.Init(servicename.ServiceName, "", "", "", configPath, nil, nil, config.ServiceNameToNamespace(mysqlconst.MysqlServiceName),
		config.ServiceNameToNamespace(redisconst.RedisServiceName),
		config.ServiceNameToNamespace(rabbitmqconst.RabbitMQServiceName),
		config.ServiceNameToNamespace(verificationconst.ServiceName),
		config.ServiceNameToNamespace(applicationconst.ServiceName))
	if err != nil {
		return xerrors.Errorf("cannot init app stub: %v", err)
	}
	err = db.Init()
	if err != nil {
		return xerrors.Errorf("cannot init database: %v", err)
	}

	return nil
}
