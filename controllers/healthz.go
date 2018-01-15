package controllers

import (
	"github.com/astaxie/beego"
)

type HealthzController struct {
	beego.Controller
}

func (c *HealthzController) Get() {
	c.TplNames = "healthz.html"
}
