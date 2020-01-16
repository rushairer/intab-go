package services

import (
	session "gopkg.in/boj/redistore.v1"
)

var _sessionInstance *session.RediStore

//SessionID 存储Session的SessionID
const SessionID = "IntabSessionID"

//InitSession 初始化Session服务
func InitSession(conf map[string]string) {

	if _sessionInstance == nil {

		store, err := session.NewRediStore(10, "tcp", conf["host"]+":"+conf["port"], conf["password"], []byte(conf["secret"]))
		if err != nil {
			panic(err)
		}

		_sessionInstance = store
	}
}

//Session 获得Session服务实例单例
func Session() *session.RediStore {
	return _sessionInstance
}

//CloseSession 关闭Session服务
func CloseSession() {
	Session().Close()
}
