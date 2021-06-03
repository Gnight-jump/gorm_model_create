/**
mysql数据库连接，使用gorm
*/
package dbConn

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"modelgenerator/conf"
	"sync"
)

var mysqlOnce sync.Once
var db *gorm.DB
var err error

func GetConnect() *gorm.DB {
	if db != nil {
		return db
	}

	mysqlOnce.Do(func() {
		baseConf := conf.Tconf.Dbase
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			baseConf.User, baseConf.Password, baseConf.Host, baseConf.Port, baseConf.DBName)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	})
	if err != nil {
		fmt.Println("[LOG_Mysql_Connect_Error] ", err)
		return nil
	}
	return db
}
