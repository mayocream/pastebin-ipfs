package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

const metadataFileName = "__metadata.json"

// UploadResp ...
type UploadResp struct {
	Cid string `json:"cid"`
}

// Object object
type Object struct {
	Name     string `json:"name"`
	MIMEType string `json:"mime_type"`
	Size     int64  `json:"size"`
}

// Metadata metadata.json
type Metadata struct {
	Author           string    `json:"author"`
	CreatedAt        time.Time `json:"created_at"`
	EncryptAlgorithm string    `json:"encrypt_algorithm"`
	Objects          []Object  `json:"objects"`
}

func (m *Metadata) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (s *Server) handleUpload(c *fiber.Ctx) error {
	blobs := make([]*file, 0)
	form, err := c.MultipartForm()
	if err != nil {
		return fiber.ErrBadRequest
	}

	meta := &Metadata{
		Author:           c.FormValue("author", "anonymous"),
		EncryptAlgorithm: c.FormValue("encrypt_algorithm", "none"),
		CreatedAt:        time.Now(),
	}

	for _, files := range form.File {
		for _, f := range files {
			if f.Size == 0 {
				continue
			}
			fr, err := f.Open()
			if err != nil {
				return fiber.ErrBadRequest
			}

			size, _ := multipartFileSize(fr)
			if size == 0 {
				// fallback
				size = f.Size
			}

			defer fr.Close()
			blobs = append(blobs, &file{
				Name:     f.Filename,
				MIMEType: mediaTypeOrDefault(f.Header),
				Reader:   fr,
				Size:     size,
			})
		}
	}

	meta.Objects = s.metadata(blobs...)

	blobs = append(blobs, &file{
		Name:   metadataFileName,
		Reader: strings.NewReader(meta.String()),
	})

	cid, err := s.creates(blobs...)
	if err != nil {
		zap.S().Errorf("create err: %s", err)
		return err
	}

	return c.Status(http.StatusCreated).JSON(cid)
}

// curl -T <file> <url>/filename
// curl -T <file> <url>/
func (s *Server) handlePut(c *fiber.Ctx) error {
	fn := c.Params("name", "plain.txt")
	if len(fn) == 0 {
		return fiber.ErrBadRequest
	}

	var err error
	fn, err = url.QueryUnescape(fn)
	if err != nil {
		return err
	}

	body := c.Body()
	if len(body) == 0 {
		return fiber.ErrBadRequest
	}

	meta := &Metadata{
		Author:           c.FormValue("author", "anonymous"),
		EncryptAlgorithm: c.FormValue("encrypt_algorithm", "none"),
		CreatedAt:        time.Now(),
	}

	blob := &file{
		Name:     fn,
		MIMEType: "plain/text",
		Reader:   bytes.NewReader(body),
		Size:     int64(len(body)),
	}

	meta.Objects = s.metadata(blob)

	cid, err := s.creates(blob, &file{
		Name:   metadataFileName,
		Reader: strings.NewReader(meta.String()),
	})

	if err != nil {
		zap.S().Errorf("creates err: %s", err)
		return err
	}

	ph := fmt.Sprintf("%s/%s/%s", c.Hostname(), cid, fn)
	c.Status(http.StatusCreated).SendString(ph)
	return nil
}

func (s *Server) handleCat(c *fiber.Ctx) error {
	cid := c.Params("cid")
	if len(cid) == 0 {
		return fiber.ErrBadRequest
	}

	// ok, err := s.idx.Exist(cid)
	// if err != nil {
	// 	zap.S().Errorf("idx get err: %s", err)
	// 	return err
	// }

	// if !ok {
	// 	return fiber.ErrBadRequest
	// }

	filename, err := url.QueryUnescape(c.Params("file"))
	if err != nil {
		return err
	}

	path := filepath.Join(cid, filename)
	src, err := s.ipc.CatStream(path)
	if err != nil {
		zap.S().Errorf("ipfs cat err: %s", err)
		return fiber.ErrInternalServerError
	}

	c.SendStream(src)
	return nil
}

// ref: https://medium.com/dtoebe/how-to-get-a-multipart-file-size-in-golang-3ab4ab4c3e3
func multipartFileSize(fr multipart.File) (int64, error) {
	switch t := fr.(type) {
	case *os.File:
		f, err := t.Stat()
		if err != nil {
			return 0, err
		}
		return f.Size(), nil
	default:
		sr, err := fr.Seek(0, 0)
		if err != nil {
			return 0, err
		}
		return sr, nil
	}
}
