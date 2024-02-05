package main

import (
	"cmd01/internal/cmd"
	_ "cmd01/internal/packed"
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gbuild"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	g.SetDebug(true)
	marshaledBytes, _ := gjson.MarshalIndent(gbuild.Info(), " ", " ")

	fmt.Println(string(marshaledBytes))

	cmd.Main.Run(gctx.GetInitCtx())
}
