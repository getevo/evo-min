package main

import (
	"github.com/getevo/evo-min"
	"github.com/getevo/evo-min/lib/log"
	"github.com/getevo/evo-min/lib/settings"
)

func main() {
	evo.Setup()

	var db = evo.GetDBO()
	_ = db

	log.Error("Application has been started", "http", settings.Get("HTTP.WriteTimeout").String(), "bool", true)
	log.SetLevel(log.DebugLevel)
	evo.Run()
}
