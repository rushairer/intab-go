package controllers

import (
	"encoding/json"
	m "intab-core/models"
	repo "intab-core/repositories"
	"log"
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/rushairer/ago"
)

//HistoryController 历史控制器
type HistoryController struct {
	BaseController
}

//HistoryListResult 显示文档返回结果类
type HistoryListResult struct {
	ago.Result
	Data *[]m.History `json:"data"`
}

//GetList 提交各种操作
func (c *HistoryController) GetList(returnBytes []byte) []byte {
	var result []byte

	key := c.Session.Keys["ch"].(string)
	uid := c.Session.Keys["uid"].(string)
	documentRepository := repo.NewDocumentRepositoryWithKey(key)
	document := documentRepository.Document
	if document.ID > 0 {

		iDid := document.ID
		iUID, _ := strconv.Atoi(uid)
		//TODO: 用户权限的判断
		log.Println(iUID)

		js, err := simplejson.NewJson([]byte(c.Msg.Data))

		if err != nil {
			log.Println(err)
			result, _ = json.Marshal(m.ResultDocument404)
		} else {

			start := js.Get("start").MustInt()
			limit := js.Get("limit").MustInt()

			historyRepository := repo.InitHistoryRepository()
			list := historyRepository.HistoryListWithDocumentIDAndLimit(iDid, start, limit)

			result, _ = json.Marshal(HistoryListResult{m.ResultDocument200, list})
		}
	} else {
		result, _ = json.Marshal(m.ResultDocument404)
	}
	returnBytes = result
	return returnBytes
}
