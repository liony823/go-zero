package main

import (
	"github.com/liony823/go-zero/core/load"
	"github.com/liony823/go-zero/core/logx"
	"github.com/liony823/go-zero/tools/goctl/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
