package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"windRise/router"
)

func main() {
	var ctx = gctx.New()
	cmd.Run(ctx)
}

var (
	cmd = gcmd.Command{
		Name:  "WindRise",
		Usage: "WindRise",
		Brief: "WindRise v1.0",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse) //默认的错误处理
				router.R.BindController(ctx, group)
			})
			s.Run()
			return nil
		},
	}
)