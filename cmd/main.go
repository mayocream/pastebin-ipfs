package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mayocream/pastebin-ipfs/pkg/index"
	"github.com/mayocream/pastebin-ipfs/pkg/ipfs"
	"github.com/mayocream/pastebin-ipfs/server"
	"go.uber.org/zap"
)

var (
	addr           = flag.String("addr", ":3939", "HTTP listen")
	ipfsAddr       = flag.String("ipfs", "127.0.0.1:5001", "IPFS address")
	debug          = flag.Bool("debug", false, "Debug mode")
	logLevel       = flag.String("log-level", "info", "Log level")
	cacheStorePath = flag.String("cache-store-path", ".pastebin", "Cache store path")
)

func main() {
	flag.Parse()
	initLogger()

	var err error
	app := server.App{
		Addr: *addr,
	}

	idx, err := index.NewIndex(*cacheStorePath)
	if err != nil {
		log.Fatal(err)
	}
	app.Index = idx

	app.IPFSClient, err = ipfs.NewClient(*ipfsAddr)
	if err != nil {
		log.Fatal(err)
	}

	if app.IPFSClient.Ping() != nil {
		log.Panic("ipfs unavailble")
	}

	server.Start(app)
}

func initLogger() {
	conf := zap.NewProductionConfig()
	conf.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	conf.Encoding = "console"

	conf.Level.UnmarshalText([]byte(*logLevel))
	logger, err := conf.Build()
	if err != nil {
		panic(fmt.Sprintf("who watches the watchmen (%v)?", err))
	}

	zap.ReplaceGlobals(logger)
}
