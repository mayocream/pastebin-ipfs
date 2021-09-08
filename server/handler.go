package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/mayocream/pastebin-ipfs/pkg/ipfs"
)

// UploadResp ...
type UploadResp struct {
	Cid string `json:"cid"`
}

// Metadata metadata
type Metadata struct {
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	// TODO files dag
}

func (m *Metadata) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (s *Server) handleUpload(c *fiber.Ctx) error {
	blobs := make([]*ipfs.File, 0)
	form, err := c.MultipartForm()
	if err != nil {
		return fiber.ErrBadRequest
	}

	for _, files := range form.File {
		for _, file := range files {
			if file.Size == 0 {
				continue
			}
			fr, err := file.Open()
			if err != nil {
				return fiber.ErrBadRequest
			}
			defer fr.Close()

			blobs = append(blobs, &ipfs.File{
				Name:   file.Filename,
				Reader: fr,
			})
		}
	}

	meta := &Metadata{
		Author:    c.FormValue("author", "unknown"),
		CreatedAt: time.Now(),
	}

	blobs = append(blobs, &ipfs.File{
		Name:   "metadata",
		Reader: strings.NewReader(meta.String()),
	})

	res, err := s.ipc.Add(blobs...)
	if err != nil {
		return err
	}

	for _, obj := range res.Objects {
		if err := s.idx.SetExist(obj.Hash); err != nil {
			return err
		}
	}

	return c.Status(http.StatusCreated).JSON(res)
}

// curl -T <file> <url>/filename
// curl -T <file> <url>/
func (s *Server) handlePut(c *fiber.Ctx) error {
	fn := c.Params("name")
	if len(fn) == 0 {
		return fiber.ErrBadRequest
	}

	body := c.Body()
	if len(body) == 0 {
		return fiber.ErrBadRequest
	}

	res, err := s.ipc.Add(&ipfs.File{
		Name:   fn,
		Reader: bytes.NewReader(body),
	})
	if err != nil {
		return err
	}

	for _, obj := range res.Objects {
		if err := s.idx.SetExist(obj.Hash); err != nil {
			return err
		}
	}

	ph := fmt.Sprintf("%s/%s/%s", c.Hostname(), res.Cid, fn)
	c.Status(http.StatusCreated).SendString(ph)
	return nil
}

func (s *Server) handleText(c *fiber.Ctx) error {
	fn := c.Params("name", "plain.txt")

	body := c.Body()
	if len(body) == 0 {
		return fiber.ErrBadRequest
	}

	res, err := s.ipc.Add(&ipfs.File{
		Name:   fn,
		Reader: bytes.NewReader(body),
	})
	if err != nil {
		return err
	}

	for _, obj := range res.Objects {
		if err := s.idx.SetExist(obj.Hash); err != nil {
			return err
		}
	}

	ph := fmt.Sprintf("%s/%s", res.Cid, fn)
	c.Status(http.StatusCreated).SendString(ph)
	return nil
}

func (s *Server) handleCat(c *fiber.Ctx) error {
	cid := c.Params("cid")
	if len(cid) == 0 {
		return fiber.ErrBadRequest
	}

	ok, err := s.idx.Exist(cid)
	if err != nil {
		return err
	}

	if !ok {
		return fiber.ErrBadRequest
	}

	src, err := s.ipc.CatStream(cid)
	if err != nil {
		log.Println("cat cid err: ", err)
		return fiber.ErrInternalServerError
	}

	c.SendStream(src)
	return nil
}
