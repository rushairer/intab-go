package models

import (
	//mysql driver
	_ "github.com/go-sql-driver/mysql"

	srv "intab-core/services"
	"time"
)

const (
	//ResourcePermissionRead 资源只读权限
	ResourcePermissionRead = 0

	//ResourcePermissionWrite 资源写权限
	ResourcePermissionWrite = 1
)

//Resource 资源类
type Resource struct {
	ID         int `gorm:"primary_key" json:"id"`
	Document   Document
	DocumentID int
	User       User
	UserID     int
	Permission int        `gorm:"not null"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index" json:"deleted_at"`
}

//NewResource 创建资源类实例
func NewResource(document Document, user User, permission int) *Resource {
	return &Resource{Document: document, User: user, Permission: permission}
}

/*
//NewResourceWithID 通过rid创建资源类实例
func NewResourceWithID(id int) *Resource {
	var resource Resource
	srv.DB().First(&resource, id)
	return &resource
}
*/

//Create 执行数据库创建命令
func (res *Resource) Create() (created bool) {
	if srv.DB().NewRecord(res) {
		srv.DB().Create(&res)
		return true
	}
	return false
}

//ListWithUserIDAndLimit 通过UserID和分页获取资源列表
func (res *Resource) ListWithUserIDAndLimit(userID int, start int, limit int) *[]Resource {
	resource := []Resource{}
	srv.DB().Limit(limit).Order("updated_at desc").Where(&Resource{UserID: userID}).Find(&resource)

	return &resource
}

//GetOneWithDocumentIDAndUserID 通过DocumnetID和UserID获得资源
func (res *Resource) GetOneWithDocumentIDAndUserID(documentID int, userID int) *Resource {
	resource := Resource{}
	srv.DB().Where(&Resource{DocumentID: documentID, UserID: userID}).First(&resource)
	return &resource
}
