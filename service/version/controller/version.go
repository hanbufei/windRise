package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"windRise/service/version/api"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Hello(ctx context.Context, req *api.Req) (res *api.Res, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("WindRise v1.0")
	return
}
