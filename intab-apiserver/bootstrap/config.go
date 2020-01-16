package bootstrap

//Config 配置
type Config struct {
	Key  string `default:"app secret key"`
	Addr string `default:":8080"`

	DB struct {
		Driver string
		DSN    string
	}

	Redis struct {
		Host     string
		Port     string
		Password string
		Secret   string
	}

	OAuth2 struct {
		ClientID          string
		ClientSecret      string
		ClientDomain      string
		WechatAPIURL      string
		WechatCallbackURL string
		WechatUserInfoURL string
	}

	NSQ struct {
		Host string
		Port string
	}
}
