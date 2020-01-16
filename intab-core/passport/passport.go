package passport

import (
	oauth2 "intab-core/oauth2"
	repo "intab-core/repositories"
	srv "intab-core/services"
	"log"
	"net/http"
)

//Passport Passport类
type Passport struct {
	AccountRepository *repo.AccountRepository
}

//PassportSessionMaxAge 存储Passport信息的过期时间
const PassportSessionMaxAge = 86400 * 7

//InitPassport 初始化一个空Passport实例
func InitPassport() *Passport {
	return &Passport{
		AccountRepository: repo.InitAccountRepository(),
	}
}

//Store 保存Passport
func (pp *Passport) Store(r *http.Request, w http.ResponseWriter) {
	session, _ := srv.Session().Get(r, srv.SessionID)

	session.Values["uid"] = pp.AccountRepository.User.ID
	session.Options.MaxAge = PassportSessionMaxAge
	session.Save(r, w)
}

//Remove 注销Passport
func (pp *Passport) Remove(r *http.Request, w http.ResponseWriter) {
	session, _ := srv.Session().Get(r, srv.SessionID)

	session.Options.MaxAge = -1
	session.Save(r, w)

	pp.AccountRepository = nil
}

//Get 读取Passport
func (pp *Passport) Get(r *http.Request, w http.ResponseWriter) bool {
	session, _ := srv.Session().Get(r, srv.SessionID)

	if uid := session.Values["uid"]; uid != nil {
		pp.AccountRepository = repo.NewAccountRepositoryWithID(uid.(int))
		return true
	}
	return false
}

//GetAccessTokenString 生成当前用户的Token字符串
func (pp *Passport) GetAccessTokenString() string {
	ti, err := oauth2.GenerateAccessTokenWithUserID(pp.AccountRepository.User.ID, "")

	if err != nil {
		log.Println(err)
	}
	//log.Println("GetAccessExpiresIn:", ti.GetAccessExpiresIn())
	return ti.GetAccess()
}
