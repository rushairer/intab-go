package controllers

import (
	pp "intab-core/passport"
	repo "intab-core/repositories"
	"net/http"

	valid "github.com/asaskevich/govalidator"
)

//AuthController Auth页面控制器
type AuthController struct {
	BaseController
}

//Login 登录页
func (c *AuthController) Login() {
	flash := c.Session.Flashes("loginFailed")
	c.SaveSession()
	c.layout("single.html")
	c.renderView("login.html", ViewData{
		"flash": flash,
	})
}

//Logout 注销
func (c *AuthController) Logout() {
	c.Passport.Remove(c.Request, c.ResponseWriter)
	http.Redirect(c.ResponseWriter, c.Request, "/dashboard", 302)
}

//PostLogin 登录
func (c *AuthController) PostLogin() {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	var accountRepository *repo.AccountRepository
	var err error
	if valid.IsEmail(username) {
		accountRepository, err = repo.NewAccountRepository(username, username, "", password)
	} else {
		accountRepository, err = repo.NewAccountRepository(username, "", username, password)
	}

	if err == nil {
		err = accountRepository.Login()

		if err == nil {
			passport := pp.Passport{}
			passport.AccountRepository = accountRepository
			passport.Store(c.Request, c.ResponseWriter)

			http.Redirect(c.ResponseWriter, c.Request, "/dashboard", 302)
		} else {
			c.loginFailed()
		}
	} else {
		c.loginFailed()
	}
}

func (c *AuthController) loginFailed() {
	c.Session.AddFlash("Login Failed", "loginFailed")
	c.SaveSession()
	http.Redirect(c.ResponseWriter, c.Request, "/login", 302)
}
