package main

import (
	_ "cmd01/internal/packed"
	"fmt"

	"github.com/gogf/gf/v2/os/gbuild"
	"github.com/gogf/gf/v2/os/gctx"

	"cmd01/internal/cmd"
)

func main() {
	fmt.Println(gbuild.Info())
	cmd.Main.Run(gctx.GetInitCtx())
}
