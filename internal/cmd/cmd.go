package cmd

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			fmt.Println("end")
			// gproc.ShellRun(ctx, `go run wx.go`)
			// gproc.ShellRun(ctx, `go run cmd/web.go`)
			return
		},
	}
)
