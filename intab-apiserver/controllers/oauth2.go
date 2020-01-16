package controllers

import (
	oauth2 "intab-core/oauth2"
	"net/http"

	"github.com/rushairer/ago"
)

//OAuth2Controller OAuth2控制器
type OAuth2Controller struct {
	ago.Controller
}

//Token Token相关
func (c *OAuth2Controller) Token() {

	//c.Request.ParseMultipartForm(32 << 20)
	//log.Println(c.Request.PostForm)

	err := oauth2.Server().HandleTokenRequest(c.ResponseWriter, c.Request)
	if err != nil {
		http.Error(c.ResponseWriter, err.Error(), http.StatusInternalServerError)
	}

}
