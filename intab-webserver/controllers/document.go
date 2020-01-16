package controllers

import (
	"fmt"
	r "intab-core/repositories"
	"log"
	"net/http"
	"time"
)

//DocumentDefaultName 文档的默认名称
const DocumentDefaultName = "无标题"

//DocumentController 文档页面控制器
type DocumentController struct {
	BaseController
}

//New 新建文档
func (c *DocumentController) New() {
	uid := c.Passport.AccountRepository.User.ID
	documentRepository := r.NewDocumentRepository(uid, DocumentDefaultName)
	if documentRepository.CreateDocument() {

		resourceRepository := r.NewResourceRepositoryPermissionWrite(*documentRepository.Document, *c.Passport.AccountRepository.User)
		if resourceRepository.CreateResource() {
			c.Session.AddFlash("Create a new document failed.", "newDocumentFailed")
			c.SaveSession()
		}
		http.Redirect(c.ResponseWriter, c.Request, fmt.Sprintf("/document/%s", documentRepository.Document.Key), 302)
	} else {
		c.Session.AddFlash("Create a new document failed.", "newDocumentFailed")
		c.SaveSession()
	}
}

//GetOne 展示一个文档
func (c *DocumentController) GetOne() {
	//TODO: 对于微信未授权用户，7日后访问要求关注公众号
	key, _ := c.Params["key"]
	documentRepository := r.NewDocumentRepositoryWithKey(key)
	document := documentRepository.Document
	//sharekey := documentRepository.Sharekey
    tokenString := c.Passport.GetAccessTokenString()

    currentTime := time.Now();

	if c.Loginned {
		resourceRepository := r.InitResourceRepository()
		resource := resourceRepository.GetResourceWithDocumentIDAndUserID(document.ID, c.Passport.AccountRepository.User.ID)

		if document.Public == 1 || resource.ID > 0 {
			//登录 公开或有Resource信息

			c.layout("document.html")
			c.renderView("document.html", ViewData{
				"document": document,
				//"sharekey":   sharekey,
				"permission": resource.Permission,
                "currentTime": currentTime,
                "token": tokenString,
                "ch": key,

			})
		} else {
			c.showNoPermission(document.Key)
		}
	} else {
		if document.Public == 0 {
			//未登录且私有
			c.showNoPermission(document.Key)
		} else {
			//未登录且公开

			c.layout("document.html")
			c.renderView("document.html", ViewData{
				"document": document,
				//"sharekey":   nil,
				"permission": 0,
                "currentTime": currentTime,
                "token": tokenString,
                "ch": key,
			})
		}
	}
}

//RequestAccess 请求文档权限
func (c *DocumentController) RequestAccess() {
	//TODO: 未登录访问被踢到登录页，登录成功后跳转回来
	key, _ := c.Params["key"]
	documentRepository := r.NewDocumentRepositoryWithKey(key)
	document := documentRepository.Document
	if document.ID > 0 {
		if c.Passport.AccountRepository.User.ID == document.UserID {
			//自己不需要搞事情
		} else {
			err := documentRepository.NewAccessRequest(c.Passport.AccountRepository.User.ID, document.ID, document.UserID)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("New Access Request!")
			}
		}
	} else {
		//文档不存在
	}
}

//AccessRequestList 权限请求列表
func (c *DocumentController) AccessRequestList() {
	documentRepository := r.InitDocumentRepository()
	list, err := documentRepository.AccessRequestListWithUserID(c.Passport.AccountRepository.User.ID)

	if err != nil {

	} else {
		log.Println(list)
	}
}

func (c *DocumentController) showNoPermission(key string) {
	c.layout("single.html")
	c.renderView("document_nopermission.html", ViewData{
		"key": key,
	})
}
