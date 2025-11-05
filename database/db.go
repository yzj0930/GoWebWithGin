package database

import (
	"fmt"
	"time"

	"github.com/yzj0930/GoWebWithGin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbConfig := config.GlobalConfig.Database
	fmt.Printf("数据库配置: 用户名=%s, 密码=%s, 主机=%s, 端口=%d, 数据库名=%s\n",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName), // data source name
		DefaultStringSize:         256,                                                                                                                                                           // default size for string fields
		DisableDatetimePrecision:  true,                                                                                                                                                          // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                                                                                                          // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                                                                                                          // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                                                                                                         // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能

	sqlDB, err := db.DB()

	if err != nil {
		panic("failed to get database, err: " + err.Error())
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
}
