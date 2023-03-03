package evo

import (
	"github.com/getevo/evo-min/lib/generic"
	"github.com/getevo/evo-min/lib/settings"
	"github.com/gofiber/fiber/v2"
	"log"
)

var (
	//Public
	app *fiber.App

	StatusCodePages = map[int]string{}
	Any             func(request *Request) error
	//private
	statics [][2]string
)
var http = HTTPConfig{}

// Setup setup the EVO app
func Setup() {
	var err = settings.Init()
	if err != nil {
		log.Fatal(err)
	}
	settings.Register("HTTP", &http)
	var fiberConfig = fiber.Config{}
	generic.Parse(http).Cast(&fiberConfig)
	app = fiber.New(fiberConfig)

	/*	if config.Server.Debug {
			fmt.Println("Enabled Logger")
			app.Use(logger.New())
			if config.Server.Recover {
				app.Use(recovermd.New())
			}

			//app.Use("/swagger", swagger.Handler) // default
		} else {
			if config.Server.Recover {
				app.Use(recovermd.New())
			}
		}*/

}

// Run start EVO Server
func Run() {
	if Any != nil {
		app.Use(func(ctx *fiber.Ctx) error {
			r := Upgrade(ctx)
			if err := Any(r); err != nil {
				return err
			}
			return nil
		})
	} else {
		// Last middleware to match anything
		app.Use(func(c *fiber.Ctx) error {
			c.SendStatus(404)
			return nil
		})
	}

	var err error
	/*	if config.Server.HTTPS {
			cer, err := tls.LoadX509KeyPair(GuessPath(config.Server.Cert), GuessPath(config.Server.Key))
			if err != nil {
				log.Fatal(err)
			}
			//err = app.Listen(config.Server.Host+":"+config.Server.Port, &tls.Config{Certificates: []tls.Certificate{cer}})
			ln, _ := net.Listen("tcp", config.Server.Host+":"+config.Server.Port)
			ln = tls.NewListener(ln, &tls.Config{Certificates: []tls.Certificate{cer}})
			err = app.Listen(config.Server.Host + ":" + config.Server.Port)
		} else {

		}*/
	err = app.Listen(http.Host + ":" + http.Port)

	log.Fatal(err)
}

// GetFiber return fiber instance
func GetFiber() *fiber.App {
	return app
}
