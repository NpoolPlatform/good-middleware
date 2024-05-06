package testinit

import (
	"context"
	"fmt"
	"path"
	"runtime"

	"github.com/NpoolPlatform/go-service-framework/pkg/app"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"

	mysqlconst "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	rabbitmqconst "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/const"
	redisconst "github.com/NpoolPlatform/go-service-framework/pkg/redis/const"
	"github.com/NpoolPlatform/good-middleware/pkg/migrator"
)

func Init() error {
	_, myPath, _, ok := runtime.Caller(0)
	if !ok {
		return wlog.Errorf("cannot get source file path")
	}

	appName := path.Base(path.Dir(path.Dir(path.Dir(myPath))))
	configPath := fmt.Sprintf("%s/../../cmd/%v", path.Dir(myPath), appName)

	err := app.Init(
		servicename.ServiceName,
		"",
		"",
		"",
		configPath,
		nil,
		nil,
		mysqlconst.MysqlServiceName,
		rabbitmqconst.RabbitMQServiceName,
		redisconst.RedisServiceName,
	)
	if err != nil {
		return wlog.Errorf("cannot init app stub: %v", err)
	}
	if err := migrator.Migrate(context.Background()); err != nil {
		panic(wlog.Errorf("fail migrate db: %v", err))
	}
	if err := db.Init(); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}
