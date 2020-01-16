package bootstrap

import (
	oauth2 "intab-core/oauth2"
	srv "intab-core/services"
	"log"
	"os"

	"github.com/facebookgo/inject"
	"github.com/jinzhu/configor"
	"github.com/rushairer/ago"
)

//App 网站程序
type App struct {
	Config *Config  `inject:""`
	Ago    *ago.Ago `inject:""`
}

var _appInstance *App

//GetApp 获得网站程序实例单例
func GetApp() *App {
	if _appInstance == nil {

		var g inject.Graph
		var app App

		err := g.Provide(
			&inject.Object{Value: &app},
		)

		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		if err := g.Populate(); err != nil {
			log.Println(err)
			os.Exit(1)
		}

		app.loadConfig()
		app.initDB()
		app.initSession()
		app.initOAuth2()
		app.initRedis()

		_appInstance = &app
	}
	return _appInstance
}

//Run 运行
func (app *App) Run() {
	app.Ago.Run(app.Config.Addr)
}

func (app *App) loadConfig() {
	configor.Load(app.Config, "config/config.yaml")
}

func (app *App) initDB() {
	srv.InitDB(map[string]string{
		"driver": app.Config.DB.Driver,
		"dsn":    app.Config.DB.DSN,
	})
}

func (app *App) initSession() {
	srv.InitSession(map[string]string{
		"host":     app.Config.Redis.Host,
		"port":     app.Config.Redis.Port,
		"password": app.Config.Redis.Password,
		"secret":   app.Config.Redis.Secret,
	})
}

func (app *App) initOAuth2() {
	oauth2.InitServer(map[string]string{
		"redis_host":          app.Config.Redis.Host,
		"redis_port":          app.Config.Redis.Port,
		"redis_password":      app.Config.Redis.Password,
		"oauth2_clientid":     app.Config.OAuth2.ClientID,
		"oauth2_clientsecret": app.Config.OAuth2.ClientSecret,
		"oauth2_clientdomain": app.Config.OAuth2.ClientDomain,
	})
}

func (app *App) initRedis() {
	srv.InitRedis(map[string]string{
		"host":     app.Config.Redis.Host,
		"port":     app.Config.Redis.Port,
		"password": app.Config.Redis.Password,
		"secret":   app.Config.Redis.Secret,
	})
}
