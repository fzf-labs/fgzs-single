package db

import (
	"database/sql"
	"fgzs-single/pkg/db/plugin/otelgorm"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewGorm(cfg *MysqlConfig) *gorm.DB {
	sqlDB, err := sql.Open("mysql", cfg.DataSourceName)
	if err != nil {
		logx.Errorf("open mysql failed! err: %+v", err)
	}
	// set for db connection
	// 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)
	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifeTime)
	gormConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	}
	if cfg.ShowLog {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gormConfig)
	if err != nil {
		logx.Errorf("database connection failed!  err: %+v", err)
	}
	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	err = db.Use(otelgorm.NewOtelPlugin())
	if err != nil {
		logx.Errorf("OpenTracing new failed!  err: %+v", err)
	}
	return db
}
