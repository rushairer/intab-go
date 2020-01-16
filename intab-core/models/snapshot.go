package models

import (
	//mysql driver
	_ "github.com/go-sql-driver/mysql"

	//"github.com/jinzhu/gorm"
	"time"
)

//Snapshot 快照类
type Snapshot struct {
	//gorm.Model
	ID         int `gorm:"primary_key"`
	User       User
	UserID     int
	DocumentID int
	HistoryID  int
	CreatedAt  time.Time
	DataJSON   string `gorm:"type:mediumtext"`
}
