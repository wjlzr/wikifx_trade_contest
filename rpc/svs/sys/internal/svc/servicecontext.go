package svc

import (
	"log"
	"os"
	"time"
	"wikifx-trade-contest/rpc/model/responseerror"
	"wikifx-trade-contest/rpc/model/user"
	"wikifx-trade-contest/rpc/svs/sys/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config     config.Config
	Eloquent   *gorm.DB
	UserModel  *user.SysUserModel
	ErrorModel *responseerror.SysErrorModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	// mysql-orm集成gorm
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)

	Eloquent, _ := gorm.Open(mysql.Open(c.Mysql.Dns), &gorm.Config{Logger: newLogger})
	return &ServiceContext{
		Config:     c,
		Eloquent:   Eloquent,
		UserModel:  user.NewSysUserModel(Eloquent),
		ErrorModel: responseerror.NewSysUserModel(Eloquent),
	}
}
