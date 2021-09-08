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

	"github.com/mayocream/pastebin-ipfs/pkg/index"
	"github.com/mayocream/pastebin-ipfs/pkg/ipfs"
)

// App config
type App struct {
	Addr       string
	IPFSClient *ipfs.Client
	Index      *index.Index
}

// Start start http server
func Start(conf App) {
	app := fiber.New(fiber.Config{
		BodyLimit: 100 << 20,
	})

	// register middlewares
	app.Use(recover.New())
	app.Use(cache.New())
	app.Use(etag.New())
	app.Use(cors.New())
	app.Use(compress.New())
	limiter.ConfigDefault.Next = func(c *fiber.Ctx) bool {
		return c.IP() == "127.0.0.1"
	}
	app.Use(limiter.New(limiter.Config{
		Max: 20,
	}))

	ipfsClient = conf.IPFSClient
	registerRoutes(app)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	go gracefulShutdown(ctx, app)

	log.Println("server listen: ", conf.Addr)
	if err := app.Listen(conf.Addr); err != nil {
		log.Fatal(err)
	}

	log.Println("Running cleanup tasks...")
}

func gracefulShutdown(ctx context.Context, app *fiber.App) {
	<-ctx.Done()
	log.Println("Gracefully shutting down...")
	app.Shutdown()
}

func registerRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Post("/text/:name", handleText)
	v1.Post("/upload", handleUpload)

	app.Post("/", handleText)
	app.Put("/:name", handlePut)

	v1.Get("/file/:cid", handleCat)
}
