package repositories

import (
	m "intab-core/models"
	srv "intab-core/services"
)

//ResourceRepository 资源处理类
type ResourceRepository struct {
	Resource *m.Resource
}

//InitResourceRepository 初始化一个空ResourceRepository实例
func InitResourceRepository() *ResourceRepository {
	return &ResourceRepository{Resource: &m.Resource{}}
}

/*
//NewResourceRepositoryWithID 通过ID创建新资源处理类实例
func NewResourceRepositoryWithID(rid int) *ResourceRepository {
	resource := m.NewResourceWithID(rid)
	resourceRepository := &ResourceRepository{Resource: resource}
	srv.DB().Model(&resourceRepository.Resource).Related(&resourceRepository.Resource.Document)
	srv.DB().Model(&resourceRepository.Resource).Related(&resourceRepository.Resource.User)
	return resourceRepository
}
*/

//NewResourceRepository 创建新资源处理类实例
func NewResourceRepository(document m.Document, user m.User, permission int) *ResourceRepository {
	resource := m.NewResource(document, user, permission)
	resourceRepository := &ResourceRepository{Resource: resource}
	srv.DB().Model(&resourceRepository.Resource).Related(&resourceRepository.Resource.Document)
	srv.DB().Model(&resourceRepository.Resource).Related(&resourceRepository.Resource.User)
	return resourceRepository
}

//NewResourceRepositoryPermissionWrite 创建可写新资源处理类实例
func NewResourceRepositoryPermissionWrite(document m.Document, user m.User) *ResourceRepository {
	return NewResourceRepository(document, user, m.ResourcePermissionWrite)
}

//NewResourceRepositoryPermissionRead 创建只读新资源处理类实例
func NewResourceRepositoryPermissionRead(document m.Document, user m.User) *ResourceRepository {
	return NewResourceRepository(document, user, m.ResourcePermissionRead)
}

//CreateResource 执行数据库创建命令
func (dp *ResourceRepository) CreateResource() (created bool) {
	return dp.Resource.Create()
}

//ResourceListWithUserIDAndLimit 通过UserID和分页获取资源列表
func (dp *ResourceRepository) ResourceListWithUserIDAndLimit(userID int, start int, limit int) *[]m.Resource {
	resources := dp.Resource.ListWithUserIDAndLimit(userID, start, limit)

	result := []m.Resource{}

	//TODO: 存在查询性能隐患，日后优化
	for _, resource := range *resources {
		srv.DB().Model(&resource).Related(&resource.Document)
		srv.DB().Model(&resource).Related(&resource.User)
		srv.DB().Model(&resource.User).Related(&resource.User.UserDetail, "UserDetail")
		srv.DB().Model(&resource.Document).Related(&resource.Document.User)
		srv.DB().Model(&resource.Document.User).Related(&resource.Document.User.UserDetail, "UserDetail")
		result = append(result, resource)
	}

	return &result
}

//GetResourceWithDocumentIDAndUserID 通过DocumnetID和UserID获得资源
func (dp *ResourceRepository) GetResourceWithDocumentIDAndUserID(documentID int, userID int) *m.Resource {
	resource := dp.Resource.GetOneWithDocumentIDAndUserID(documentID, userID)

	return resource
}
