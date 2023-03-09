package main

import (
	"fmt"
	"github.com/getevo/evo-min"
	"github.com/getevo/evo-min/lib/cache"
	"github.com/getevo/evo-min/lib/log"
	"github.com/getevo/evo-min/lib/settings"
	"time"
)

func main() {
	evo.Setup()

	cache.Register()

	cache.Set("x", 196, 5*time.Second)
	var x int
	cache.Get("x", &x)
	fmt.Println("x1:", x)
	time.Sleep(6 * time.Second)

	fmt.Println("x2:", cache.Get("x", &x), x)
	var db = evo.GetDBO()
	_ = db
	log.Error("Application has been started", "http", settings.Get("HTTP.WriteTimeout").String(), "bool", true)
	log.SetLevel(log.DebugLevel)
	evo.Run()
}
