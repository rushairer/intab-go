package services

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/mssql"
)

var _dbInstance *gorm.DB

//InitDB 初始化数据库服务
func InitDB(conf map[string]string) {
	if _dbInstance == nil {
		db, err := gorm.Open(conf["driver"], conf["dsn"])

		db.DB().SetMaxOpenConns(100)
		db.DB().SetConnMaxLifetime(time.Minute * 5)
		db.DB().SetMaxIdleConns(0)
		db.DB().SetMaxOpenConns(5)

		if err != nil {
			log.Println("failed to connect database:")
			panic(err)
		}
		_dbInstance = db
	}
}

//DB 获得DB服务实例单例
func DB() *gorm.DB {
	return _dbInstance
}

//CloseDB 关闭数据库连接
func CloseDB() {
	DB().Close()
}
