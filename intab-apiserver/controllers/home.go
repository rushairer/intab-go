package controllers

import (
	"github.com/rushairer/ago"
)

//HomeController Home控制器
type HomeController struct {
	ago.Controller
}

//Get Get方法
func (c *HomeController) Get() {
	c.JSON(ago.Result200)
}

//Index Index方法
func (c *HomeController) Index() {
	c.JSON(ago.Result200)
}
