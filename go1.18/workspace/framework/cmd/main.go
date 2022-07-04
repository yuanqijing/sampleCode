package main

import (
	"github.com/yuanqijing/framework/pkg"
	plugin_1 "github.com/yuanqijing/framework/pkg/plugin-1"
)

func main() {

	var hello pkg.Hello = &plugin_1.Plugin{
		Name: "plugin_1",
	}

	hello.Hello()
}
