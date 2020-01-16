package controllers

import (
	m "intab-core/models"

	"github.com/rushairer/ago"
)

//DocumentShowResult 显示文档返回结果类
type DocumentShowResult struct {
	ago.Result
	Data *m.Document `json:"data"`
}

//DocumentSharekeyResult 显示文档返回结果类
type DocumentSharekeyResult struct {
	ago.Result
	Data string `json:"data"`
}

//DocumentController 文档控制器
type DocumentController struct {
	BaseController
}

/*
//CreateSharekey 创建sharekey
func (c *DocumentController) CreateSharekey() {
	documentID, _ := strconv.Atoi(c.Request.FormValue("did"))

	documentRepository := r.NewDocumentRepositoryWithID(documentID)

	var result DocumentSharekeyResult

	if c.Passport.AccountRepository.User.ID == documentRepository.Document.UserID {
		documentRepository.UpdateSharekey()

		if documentRepository.Sharekey.ID > 0 {
			result = DocumentSharekeyResult{m.ResultSharekey200, documentRepository.Sharekey.Key}
		} else {
			result = DocumentSharekeyResult{m.ResultSharekey400, ""}
		}
	} else {
		result = DocumentSharekeyResult{m.ResultSharekey403, ""}
	}
	c.JSON(result)
}

//GetOne 通过documentID获得一个文档
func (c *DocumentController) GetOne() {
	documentID, _ := strconv.Atoi(c.Params["documentID"])

	document := r.NewDocumentRepositoryWithID(documentID).Document

	var result DocumentShowResult

	if document.ID > 0 {
		result = DocumentShowResult{m.ResultDocument200, nil}
		result.Data = document
	} else {
		result = DocumentShowResult{m.ResultDocument404, nil}
	}

	c.JSON(result)
}

//Create 创建文档
func (c *DocumentController) Create() {
	uid := c.Passport.AccountRepository.User.ID

	filename := c.Request.FormValue("filename")

	var result DocumentShowResult

	if r.NewDocumentRepository(uid, filename).CreateDocument() {
		result = DocumentShowResult{m.ResultDocument201, nil}
	} else {
		result = DocumentShowResult{m.ResultDocument400, nil}
	}
	c.JSON(result)
}
*/
