package mysql

import (
	"fmt"
	"sync"

	"github.com/nkien0204/lets-go/internal/configs"
	"github.com/nkien0204/rolling-logger/rolling"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlService struct {
	Address string
	Db      *gorm.DB
}

var service *MysqlService
var once sync.Once

func GetMysqlConnection() *MysqlService {
	once.Do(func() {
		var err error
		addr := configs.GetConfigs().Db.Addr
		if service, err = initMysqlConnection(addr); err != nil {
			rolling.New().Error("initMysqlConnection failed", zap.Error(err))
			panic(1)
		}
	})
	return service
}

func initMysqlConnection(addr string) (*MysqlService, error) {
	logger := rolling.New()
	addr = fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", addr)
	db, err := gorm.Open(mysql.Open(addr))
	if err != nil {
		logger.Error("gorm.Open failed", zap.Error(err))
		return nil, err
	}

	return &MysqlService{
		Address: addr,
		Db:      db,
	}, nil
}
