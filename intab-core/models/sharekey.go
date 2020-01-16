package models

import (
	"log"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"

	srv "intab-core/services"
	"time"

	"github.com/segmentio/ksuid"
)

//Sharekey 分享密钥类
type Sharekey struct {
	ID         int    `gorm:"primary_key" json:"id"`
	Key        string `gorm:"unique"`
	Document   Document
	DocumentID int
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index" json:"deleted_at"`
}

//NewSharekeyWithKey 通过Key读取分享密钥类实例
func NewSharekeyWithKey(key string) *Sharekey {
	var sharekey Sharekey
	srv.DB().Where(&Sharekey{Key: key}).First(&sharekey)
	return &sharekey
}

//NewSharekeyWithDocumentID 通过DocumentID读取分享密钥类实例
func NewSharekeyWithDocumentID(documentID int) *Sharekey {
	var sharekey Sharekey
	srv.DB().Where(&Sharekey{DocumentID: documentID}).First(&sharekey)
	return &sharekey
}

//NewSharekeyWithDocumentIDOrCreate 创建分享密钥实例
func NewSharekeyWithDocumentIDOrCreate(documentID int) *Sharekey {
	var sharekey Sharekey
	key := ksuid.New().String()
	log.Println("DocumentID:", documentID, " key:", key)
	srv.DB().Where(Sharekey{DocumentID: documentID}).Assign(Sharekey{Key: key}).FirstOrCreate(&sharekey)

	return &sharekey
}
