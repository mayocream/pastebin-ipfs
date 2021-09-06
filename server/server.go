package server

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mayocream/pastebin-ipfs/pkg/ipfs"
)

// App config
type App struct {
	Addr       string
	IPFSClient *ipfs.Client
}

// Start start http server
func Start(conf App) {
	app := fiber.New()

	// register middlewares
	app.Use(recover.New())
	app.Use(cache.New())
	app.Use(etag.New())
	app.Use(cors.New())
	app.Use(compress.New())
	limiter.ConfigDefault.Next = func(c *fiber.Ctx) bool {
		return c.IP() == "127.0.0.1"
	}
	app.Use(limiter.New())

	ipfsClient = conf.IPFSClient
	registerRoutes(app)

	ctx := Graceful()
	go func() {
		<-ctx.Done()
		log.Println("Gracefully shutting down...")
		app.Shutdown()
	}()

	log.Println("server listen: ", conf.Addr)
	if err := app.Listen(conf.Addr); err != nil {
		log.Panic(err)
	}

	log.Println("Running cleanup tasks...")
}

// Graceful context
func Graceful() context.Context {
	ctx := context.Background()
	ctx, _ = signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	return ctx
}

func registerRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Post("/file", handleAdd)
	v1.Get("/file/:cid", handleCat)
}
