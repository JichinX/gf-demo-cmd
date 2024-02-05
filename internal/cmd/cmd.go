package cmd

import (
	"cmd01/internal/controller/hello"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:        "target",
		Description: "指定当前运行的模式，默认 web",
		Usage:       "-target=[web|wx]",
		Brief:       "ss",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			targert := parser.GetOpt("target", "web").Val()
			createServer(targert).Run()
			return
		},
	}
)

func createServer(target interface{}) *ghttp.Server {
	s := g.Server()
	if target == "web" {
		s.Group("/api/v1", func(group *ghttp.RouterGroup) {
			group.Middleware(ghttp.MiddlewareHandlerResponse)
			group.Bind(
				hello.NewV1(),
			)
		})
	} else {
		s.Group("/api/v1/wx", func(group *ghttp.RouterGroup) {
			group.Middleware(ghttp.MiddlewareHandlerResponse)
			group.Bind(
				hello.NewV1(),
			)
		})
	}
	return s
}
