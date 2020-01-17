package main

import (
	_ "blog/boot"
	_ "blog/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
