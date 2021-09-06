package server

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mayocream/pastebin-ipfs/pkg/ipfs"
)

var ipfsClient *ipfs.Client

func handleAdd(c *fiber.Ctx) error {
	// c.Request().Header.ContentLength()
	// TODO max size limit
	file := c.Body()
	if len(file) == 0 {
		return fiber.ErrBadRequest
	}

	cid, err := ipfsClient.Add(file)
	if err != nil {
		log.Println("add file err: ", err)
		return fiber.ErrInternalServerError
	}

	c.Status(http.StatusCreated).SendString(cid)
	return nil
}

func handleCat(c *fiber.Ctx) error {
	cid := c.Params("cid")
	if len(cid) == 0 {
		return fiber.ErrBadRequest
	}

	src, err := ipfsClient.CatStream(cid)
	if err != nil {
		log.Println("cat cid err: ", err)
		return fiber.ErrInternalServerError
	}

	c.SendStream(src)
	return nil
}
