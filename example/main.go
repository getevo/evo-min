package main

import (
	"fmt"
	"github.com/getevo/evo-min"
	"github.com/getevo/evo-min/lib/cache"
	"github.com/getevo/evo-min/lib/cache/drivers/redis"
	"github.com/getevo/evo-min/lib/log"
	"github.com/getevo/evo-min/lib/pubsub"
	"github.com/getevo/evo-min/lib/pubsub/drivers/kafka"
	"github.com/getevo/evo-min/lib/settings"
	"github.com/getevo/evo-min/lib/settings/database"
	"time"
)

func main() {
	evo.Setup()

	var db = evo.GetDBO()
	var data = map[string]interface{}{}
	db.Debug().Raw("SELECT * FROM services").Scan(&data)

	settings.SetDefaultDriver(database.Driver)
	cache.SetDefaultDriver(redis.Driver)

	pubsub.AddDriver(redis.Driver)
	pubsub.SetDefaultDriver(kafka.Driver)

	pubsub.Use("redis").Subscribe("test", func(topic string, message []byte, driver pubsub.Interface) {
		log.Debug("message received", "driver", driver.Name(), "topic", topic, "message", string(message))
	})

	go func() {
		for {
			pubsub.Use("redis").Publish("test", []byte(fmt.Sprint(time.Now().Unix())))
			time.Sleep(1 * time.Second)
		}
	}()

	log.Error("Application has been started", "http", settings.Get("HTTP.WriteTimeout").String(), "bool", true)
	log.SetLevel(log.DebugLevel)
	evo.Run()
}
