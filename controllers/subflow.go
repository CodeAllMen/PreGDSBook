/*
Package controllers

Author  navi 2020/2/28 - 10:10 上午
Updated navi 2020/2/28 - 10:10 上午
*/
package controllers

import (
	"github.com/MobileCPX/PreBaseLib/common"
)

type SubFlowController struct {
	common.BaseController
}

func (c *SubFlowController) Get() {
	//trackID := c.GetString("track")
	//
	//if trackID != "" {
	//	c.Data["URL"] = "http://fr.allcpx.com/start-sub?track=" + trackID
	//	c.TplName = "lp.html"
	//} else {
	//	// LP 页面存入此次点击信息，获取aff_track 表自增ID
	//	trackID := tracking.LpPageTracking(c.Ctx.Request, "http://fr.allcpx.com/aff/click", "112419-GDS")
	//	// 将trackID转为int类型，判断trackID是否为数字类型
	//	_, err := strconv.Atoi(trackID)
	//
	//	if err != nil { // 说明trackID不是int类型，不能订阅服务
	//		c.Ctx.ResponseWriter.ResponseWriter.WriteHeader(404)
	//		c.StringResult("wrong track param")
	//		c.StopRun()
	//	}
	//	// 重定向到用户标识页面
	//	c.Redirect("http://fr.allcpx.com/service/identify?track="+trackID, 302)
	//	//c.StopRun()
	//}

	c.TplName = "lp.html"
}

// 服务管理页
func (c *SubFlowController) ServiceManagement() {
	c.TplName = "help.html"
}

// 兼容性设备页
func (c *SubFlowController) Compat() {
	c.TplName = "compat.html"
}

// 信息条款页
func (c *SubFlowController) InfoPrivacy() {
	c.TplName = "privacy.html"
}

// 条款页
func (c *SubFlowController) TNC() {
	c.TplName = "tnc.html"
}
