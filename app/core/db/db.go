package db

import (
	. "stencil-go/app/core/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func New() *gorm.DB {
	dialect := GConfig.Dialect

	switch dialect {
	case "mysql":
		db, err := gorm.Open(mysql.Open(GConfig.MysqlUrl()), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单数表名
			},
			Logger: logger.Default.LogMode(logger.Silent),
		})
		sqlDB, err := db.DB()

		if GConfig.MySQL.Required {
			if err = sqlDB.Ping(); err != nil {
				panic(err.Error())
			}
		}

		sqlDB.SetMaxOpenConns(GConfig.MySQL.MaxOpenConns)
		sqlDB.SetMaxIdleConns(GConfig.MySQL.MaxIdleConns)

		return db
	case "postgre":
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  GConfig.PostgreUrl(),
			PreferSimpleProtocol: true, // 禁用隐式 prepared statement
		}), &gorm.Config{})

		if err != nil {
			panic(err)
		}
		return db
	}
	return nil
}
