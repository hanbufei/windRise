package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	version "windRise/service/version/controller"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	//version 路由，绑定url和controller
	group.Group("/version", func(group *ghttp.RouterGroup) {
		group.Bind(
			version.New(),
		)
	})
}
