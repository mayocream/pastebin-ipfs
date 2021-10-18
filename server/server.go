package server

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
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
	limiter.ConfigDefault.Next = func(c *fiber.Ctx) bool {
		return c.IP() == "127.0.0.1"
	}
	app.Use(limiter.New(limiter.Config{
		Max: 20,
        KeyGenerator: func(c *fiber.Ctx) string {
            if ip := c.GetRespHeader("CF-Connecting-IP"); ip != "" {
                return ip
            }
            return c.IP()
        },
	}))
    app.Use(func(c *fiber.Ctx) error {
        if c.Method() != fiber.MethodGet {
            return c.Next()
        }
        if err := c.Next(); err != nil {
			return err
		}
        c.Set(fiber.HeaderCacheControl, "public, max-age=604800")
        return nil
    })

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
	// API Usage
    v0 := app.Group("/api/v0")
	v0.Post("/upload", s.handleUpload)
    v0.Get("/gallery", s.handleGallery)
    v0.Get("/:cid/:file", s.handleCat)

    // Terminal Upload
    v0.Post("/", s.handlePut)
    v0.Put("/:name", s.handlePut)
    
    // Root shorter upload url. 
	// app.Post("/", s.handlePut)
	// app.Put("/:name", s.handlePut)

	app.Get("/:cid/raw/:file", s.handleCat)

    // IPFS paths
    app.Get("/ipfs/:cid/:file", s.handleCat)

    // Static files or frontend
    app.Static("/", "./dist")
    app.Get("/*", func(c *fiber.Ctx) error {
        return c.SendFile("./dist/index.html")
    })
}
