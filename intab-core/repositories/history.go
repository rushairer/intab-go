package repositories

import (
	m "intab-core/models"

	valid "github.com/asaskevich/govalidator"
)

//HistoryRepository 历史处理类
type HistoryRepository struct {
	History *m.History
}

//InitHistoryRepository 初始化一个空HistoryRepository实例
func InitHistoryRepository() *HistoryRepository {
	return &HistoryRepository{History: &m.History{}}
}

//NewHistoryRepository 创建历史处理类实例
func NewHistoryRepository(userID int, documentID int, htype int, action int, address string, content string) (*HistoryRepository, error) {
	history := m.NewHistory(userID, documentID, htype, action, address, content)

	result, err := valid.ValidateStruct(history)
	if err != nil {
		return nil, err
	} else if result {
		historyRepository := &HistoryRepository{History: history}
		return historyRepository, nil
	} else {
		return nil, nil
	}
}

//CreateHistory 执行数据库创建命令
func (hp *HistoryRepository) CreateHistory() (created bool) {
	if err := hp.History.Create(); err != nil {
		return false
	}
	return true
}

//HistoryListWithDocumentIDAndLimit 通过DocumentID和分页获取历史列表
func (hp *HistoryRepository) HistoryListWithDocumentIDAndLimit(documentID int, start int, limit int) *[]m.History {
	return hp.History.ListWithDocumentIDAndLimit(documentID, start, limit)
}
