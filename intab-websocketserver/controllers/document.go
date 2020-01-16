package controllers

import (
	"encoding/json"
	m "intab-core/models"
	repo "intab-core/repositories"
	"log"
	"strconv"

	"github.com/bitly/go-simplejson"
)

//DocumentController 文档控制器
type DocumentController struct {
	BaseController
}

//GetDocuemtContent 获取文档内容
func (c *DocumentController) GetDocuemtContent(returnBytes []byte) []byte {
	key := c.Session.Keys["ch"].(string)

	log.Println("Document ID: ", key)

	result := `
	{
		"meta": {
			"count": {
				"row": 7,
				"column": 2
			}
		},
		"cell": {
			"width": [180, 85],
			"height": [40, 98, 40, 50, 40, 40, 40]
		},
		"columns": [
			["周一", "周二", "周三", "周四", "周五", "周六", "周日"],
			["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
		]
	}`
	returnBytes = []byte(result)

	return returnBytes
}

//Commit 提交各种操作
func (c *DocumentController) Commit(returnBytes []byte) []byte {
	var result []byte

	key := c.Session.Keys["ch"].(string)
	uid := c.Session.Keys["uid"].(string)
	documentRepository := repo.NewDocumentRepositoryWithKey(key)
	document := documentRepository.Document
	if document.ID > 0 {

		iDid := document.ID
		iUID, _ := strconv.Atoi(uid)

		js, err := simplejson.NewJson([]byte(c.Msg.Data))

		if err != nil {
			log.Println(err)
			result, _ = json.Marshal(m.ResultCommit400)
		} else {

			address := js.Get("address").MustString()
			htype := js.Get("type").MustInt()
			action := js.Get("action").MustInt()
			content := js.Get("content").MustString()

			historyRepository, err := repo.NewHistoryRepository(iUID, iDid, htype, action, address, content)
			if err != nil {
				log.Println(err)
				result, _ = json.Marshal(m.ResultCommit400)
			} else {
				created := historyRepository.CreateHistory()
				if created {
					result, _ = json.Marshal(m.ResultCommit200)
				} else {
					result, _ = json.Marshal(m.ResultCommit400)
				}
			}
		}
	} else {
		result, _ = json.Marshal(m.ResultDocument404)
	}
	returnBytes = result
	return returnBytes
}
