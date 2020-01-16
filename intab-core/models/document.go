package models

import (
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/segmentio/ksuid"

	srv "intab-core/services"
	"time"
)

//Document 文档类
type Document struct {
	ID        int    `gorm:"primary_key" json:"id"`
	Key       string `gorm:"type:varchar(32);not null;unique"`
	User      User
	UserID    int
	Filename  string `gorm:"type:varchar(255);not null" json:"filename"`
	Public    int
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

//NewDocument 创建文档类实例
func NewDocument(userID int, filename string) *Document {
	key := ksuid.New().String()
	return &Document{Key: key, UserID: userID, Filename: filename}
}

/*
//NewDocumentWithID 通过ID创建文档类实例
func NewDocumentWithID(id int) *Document {
	var document Document
	srv.DB().First(&document, id)
	return &document
}
*/

//NewDocumentWithKey 通过Key创建文档类实例
func NewDocumentWithKey(key string) *Document {
	var document Document
	srv.DB().Where(&Document{Key: key}).First(&document)
	return &document
}

//Create 执行数据库创建命令
func (doc *Document) Create() (created bool) {
	if srv.DB().NewRecord(doc) {
		srv.DB().Create(&doc)
		return true
	}
	return false
}

//ListWithUserIDAndLimit 通过UserID和分页获取文档列表
func (doc *Document) ListWithUserIDAndLimit(userID int, start int, limit int) *[]Document {
	document := []Document{}
	srv.DB().Offset(start).Limit(limit).Order("updated_at desc").Where(&Document{UserID: userID}).Find(&document)

	return &document
}
