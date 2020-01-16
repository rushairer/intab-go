package repositories

import (
	"fmt"
	m "intab-core/models"
	srv "intab-core/services"

	"github.com/garyburd/redigo/redis"
)

//DocumentRepository 文档处理类
type DocumentRepository struct {
	Document *m.Document
	//Sharekey *m.Sharekey
}

//DocumentRepositoryAccessRequestPrefix AccessRequest的KV存储键
const DocumentRepositoryAccessRequestPrefix = "AccessRequest"

const (
	//DocumentPermissionRead 读权限
	DocumentPermissionRead = 1 << 0

	//DocumentPermissionWrite 写权限
	DocumentPermissionWrite = 1 << 1

	//DocumentPermissionDelete 删除权限
	DocumentPermissionDelete = 1 << 2
)

//InitDocumentRepository 初始化一个空DocumentRepository实例
func InitDocumentRepository() *DocumentRepository {
	return &DocumentRepository{Document: &m.Document{}}
}

//NewDocumentRepositoryWithKey 通过ID创建新文档处理类实例
func NewDocumentRepositoryWithKey(key string) *DocumentRepository {
	document := m.NewDocumentWithKey(key)
	//sharekey := m.NewSharekeyWithDocumentID(documentID)
	//return &DocumentRepository{Document: document, Sharekey: sharekey}
	return &DocumentRepository{Document: document}
}

/*
//NewDocumentRepositoryWithID 通过ID创建新文档处理类实例
func NewDocumentRepositoryWithID(documentID int) *DocumentRepository {
	document := m.NewDocumentWithID(documentID)
	//sharekey := m.NewSharekeyWithDocumentID(documentID)
	//return &DocumentRepository{Document: document, Sharekey: sharekey}
	return &DocumentRepository{Document: document}
}
*/

//NewDocumentRepository 创建新文档处理类实例
func NewDocumentRepository(userID int, filename string) *DocumentRepository {
	document := m.NewDocument(userID, filename)
	return &DocumentRepository{Document: document}
}

//CreateDocument 执行文档数据库创建命令
func (dp *DocumentRepository) CreateDocument() (created bool) {
	return dp.Document.Create()
}

//DocumentListWithUserIDAndLimit 通过UserID和分页获取文档列表
func (dp *DocumentRepository) DocumentListWithUserIDAndLimit(userID int, start int, limit int) *[]m.Document {
	return dp.Document.ListWithUserIDAndLimit(userID, start, limit)
}

//GetPermissionWithUserID 获得用户的权限
func (dp *DocumentRepository) GetPermissionWithUserID(userID int) int {
	permission := 0

	if userID == dp.Document.UserID {
		permission = permission | DocumentPermissionRead | DocumentPermissionWrite | DocumentPermissionDelete
	} else {
		resourceRepository := InitResourceRepository()
		resource := resourceRepository.GetResourceWithDocumentIDAndUserID(dp.Document.ID, userID)

		if resource.ID > 0 {
			if resource.Permission == m.ResourcePermissionWrite {
				permission = permission | DocumentPermissionWrite
			} else {
				permission = permission | DocumentPermissionRead
			}
		} else {
			if dp.Document.Public == 1 {
				permission = permission | DocumentPermissionRead
			}
		}
	}

	return permission
}

/*
//UpdateSharekey 更新分享密钥数据库
func (dp *DocumentRepository) UpdateSharekey() {
	dp.Sharekey = m.NewSharekeyWithDocumentIDOrCreate(dp.Document.ID)
	log.Println(dp.Sharekey)
}
*/

//NewAccessRequest 新增权限请求
func (dp *DocumentRepository) NewAccessRequest(userID int, documentID int, toUserID int) error {
	_, err := srv.Redis().Do("HSET", fmt.Sprintf("%s_%d", DocumentRepositoryAccessRequestPrefix, toUserID), fmt.Sprintf("%d|%d", documentID, userID), 1)
	return err
}

//AccessRequestListWithUserID 获取到达该用户的权限请求通知
func (dp *DocumentRepository) AccessRequestListWithUserID(userID int) (*map[string]string, error) {
	result, err := redis.StringMap(srv.Redis().Do("HGETALL", fmt.Sprintf("%s_%d", DocumentRepositoryAccessRequestPrefix, userID)))
	return &result, err
}

//RemoveAccessRequest 删除权限请求
func (dp *DocumentRepository) RemoveAccessRequest(userID int, documentID int, toUserID int) error {
	_, err := srv.Redis().Do("HDEL", fmt.Sprintf("%s_%d", DocumentRepositoryAccessRequestPrefix, toUserID), fmt.Sprintf("%d|%d", documentID, userID))
	return err
}
