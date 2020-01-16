package oauth2

import (
	"log"
	//"net/http"
	"fmt"
	repo "intab-core/repositories"

	"github.com/go-oauth2/redis"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

var _oauth2Instance *server.Server
var _oauth2Conf map[string]string

//InitServer 初始化Oauth2服务
func InitServer(conf map[string]string) {

	if _oauth2Instance == nil {

		_oauth2Conf = conf

		manager := manage.NewDefaultManager()
		// token memory store
		// manager.MustTokenStorage(store.NewMemoryTokenStore())

		// use redis token store
		manager.MustTokenStorage(redis.NewTokenStore(&redis.Config{
			Addr:     conf["redis_host"] + ":" + conf["redis_port"],
			Password: conf["redis_password"],
		}))

		// client memory store

		clientStore := store.NewClientStore()
		clientStore.Set(conf["oauth2_clientid"], &models.Client{
			ID:     conf["oauth2_clientid"],
			Secret: conf["oauth2_clientsecret"],
			Domain: conf["oauth2_clientdomain"],
		})
		manager.MapClientStorage(clientStore)

		srv := server.NewDefaultServer(manager)
		srv.SetAllowGetAccessRequest(true)
		srv.SetClientInfoHandler(server.ClientFormHandler)
		/*
		   srv.SetUserAuthorizationHandler( func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		       //err = errors.ErrAccessDenied
		       userID = "123"
		       return
		   })
		*/

		srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
			//err = errors.ErrAccessDenied
			accountRepository, err := repo.NewAccountRepository(username, "", "", password)
			if err != nil {
				return
			}

			err = accountRepository.Login()
			if err != nil {
				return
			}

			userID = fmt.Sprintf("%d", accountRepository.User.ID)
			return
		})

		srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
			log.Println("OAuth Internal Error:", err.Error())
			return
		})

		srv.SetResponseErrorHandler(func(re *errors.Response) {
			log.Println("OAuth Response Error:", re.Error.Error())
		})

		_oauth2Instance = srv
	}
}

//Server 获得OAuth2服务实例单例
func Server() *server.Server {
	return _oauth2Instance
}

//GenerateAccessTokenWithUserID 通过UserID生成Token
func GenerateAccessTokenWithUserID(userID int, scope string) (ti oauth2.TokenInfo, err error) {
	tgr := &oauth2.TokenGenerateRequest{
		ClientID:     _oauth2Conf["oauth2_clientid"],
		ClientSecret: _oauth2Conf["oauth2_clientsecret"],
		Scope:        scope,
		UserID:       fmt.Sprintf("%d", userID),
	}

	ti, err = _oauth2Instance.Manager.GenerateAccessToken(oauth2.PasswordCredentials, tgr)
	return
}
