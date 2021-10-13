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

// Config ...
type Config struct {
	IPFSClient *ipfs.Client
	Index      *index.Index
}

// Server server
type Server struct {
	ipc *ipfs.Client
	idx *index.Index
}

// New ...
func New(conf *Config) *Server {
	return &Server{
		ipc: conf.IPFSClient,
		idx: conf.Index,
	}
}

// Start start http server
func (s *Server) Start(addr string) {
	app := fiber.New(fiber.Config{
		BodyLimit: 1 << 20,
	})

	// register middlewares
    // middleware order matters!
	app.Use(recover.New())
	app.Use(etag.New())
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(cache.New())
	limiter.ConfigDefault.Next = func(c *fiber.Ctx) bool {
		return c.IP() == "127.0.0.1"
	}
	app.Use(limiter.New(limiter.Config{
		Max: 20,
	}))

	s.registerRoutes(app)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	go gracefulShutdown(ctx, app)

	log.Println("server listen at ", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}

	log.Println("Running cleanup tasks...")
}

func gracefulShutdown(ctx context.Context, app *fiber.App) {
	<-ctx.Done()
	log.Println("Gracefully shutting down...")
	app.Shutdown()
}

func (s *Server) registerRoutes(app *fiber.App) {
	v0 := app.Group("/api/v0")
	v0.Post("/upload", s.handleUpload)
    v0.Get("/gallery", s.handleGallery)
    v0.Get("/:cid/:file", s.handleCat)
    v0.Post("/", s.handlePut)
    v0.Put("/:name", s.handlePut)
    
	app.Post("/", s.handlePut)
	app.Put("/:name", s.handlePut)

	app.Get("/:cid/raw/:file", s.handleCat)

    app.Static("/", "./dist")
}
