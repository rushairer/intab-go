package models

import (
	//mysql driver
	_ "github.com/go-sql-driver/mysql"

	srv "intab-core/services"
	"time"
)

const (
	//HistoryTypeCell History类型为Cell
	HistoryTypeCell = 1

	//HistoryTypeRow History类型为Row
	HistoryTypeRow = 2

	//HistoryTypeColumn History类型为Column
	HistoryTypeColumn = 3

	//HistoryTypeRectangle History类型为Rectangle
	HistoryTypeRectangle = 4

	//HistoryActionInsert History的Insert动作
	HistoryActionInsert = 1

	//HistoryActionMove History的Move动作
	HistoryActionMove = 2

	//HistoryActionDelete History的Delete动作
	HistoryActionDelete = 3

	//HistoryActionHide History的Hide动作
	HistoryActionHide = 4

	//HistoryActionShow History的Show动作
	HistoryActionShow = 5

	//HistoryActionResize History的Resize动作
	HistoryActionResize = 6

	//HistoryActionAutoResize History的AutoResize动作
	HistoryActionAutoResize = 7
)

//History 历史记录类
type History struct {
	//gorm.Model
	ID         int `gorm:"primary_key" json:"id"`
	UserID     int
	DocumentID int    `json:"document_id" valid:"required"`
	Type       int    `valid:"required"`
	Action     int    `valid:"required"`
	Address    string `gorm:"type:text"`
	Content    string `gorm:"type:mediumblob"`
	CreatedAt  time.Time
}

//NewHistory 创建历史记录类实例
func NewHistory(userID int, documentID int, htype int, action int, address string, content string) *History {
	return &History{
		UserID:     userID,
		DocumentID: documentID,
		Type:       htype,
		Action:     action,
		Address:    address,
		Content:    content,
	}
}

//Create 创建历史记录
func (h *History) Create() error {
	return srv.DB().Create(&h).Error
}

//ListWithDocumentIDAndLimit 通过DocumentID和分页获取文档列表
func (h *History) ListWithDocumentIDAndLimit(documentID int, start int, limit int) *[]History {
	history := []History{}
	srv.DB().Offset(start).Limit(limit).Order("id asc").Where(&History{DocumentID: documentID}).Find(&history)

	return &history
}
