package main

import (
	"github.com/unstoppablego/framework/config"
	"github.com/unstoppablego/framework/httpapi"
	"github.com/unstoppablego/framework/logs"
	"github.com/unstoppablego/unstoppablego/controls"
)

func main() {

	config.ReadConf("./configs/") //"../../configs/"

	logs.Info("Hello world")

	controls.Api()

	httpapi.Provider().RunServer("0.0.0.0:1999", nil)
}
