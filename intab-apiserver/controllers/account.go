package controllers

import (
	repo "intab-core/repositories"

	"github.com/rushairer/ago"
)

//AccountController 账户控制器
type AccountController struct {
	ago.Controller
}

//Prepare 运行在每个Action之前
func (c *AccountController) Prepare() {
}

//Create 创建账户
func (c *AccountController) Create() {
	//TODO 短信或验证码验证
	name := c.Request.FormValue("name")
	email := c.Request.FormValue("email")
	phone := c.Request.FormValue("phone")
	password := c.Request.FormValue("password")

	accountRepository, err := repo.NewAccountRepository(name, email, phone, password)

	if err == nil {
		accountRepository.LoginOrCreate()
	}
}
