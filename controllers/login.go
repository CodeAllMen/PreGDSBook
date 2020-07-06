package controllers

import (
	"fmt"
	"github.com/MobileCPX/PreGDSBook/models"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["status"] = "0"
	c.TplName = "login.html"
}

func (c *UserController) Get() {
	switch c.Ctx.Input.Param(":method") {
	case "login":
		// if c.GetString("user") == "admin" && c.GetString("pass") == "test123" {
		// 	c.SetSession("user", c.GetString("user"))
		// }
		user := c.GetString("user")
		if user != "" {
			userID := models.CheckUser(user)
			if userID != "" { // 登录成功
				fmt.Println("=========================> userid:", userID)
				c.SetSession("user", userID)
				c.Redirect("/", 302)
				return
			}
		}
	}
	c.Data["status"] = "1" // 登录失败
	c.TplName = "login.html"
}

func (c *UserController) Post() {
	switch c.Ctx.Input.Param(":method") {
	case "login":
		// if c.GetString("user") == "admin" && c.GetString("pass") == "test123" {
		// 	c.SetSession("user", c.GetString("user"))
		// }
		user := c.GetString("user")
		if user != "" {
			userID := models.CheckUser(user)
			if userID != "" { // 登录成功
				fmt.Println("=========================> userid:", userID)
				c.SetSession("user", userID)
				c.Redirect("/", 302)
				return
			}
		}
	}
	c.Data["status"] = "1" // 登录失败
	c.TplName = "login.html"
}
