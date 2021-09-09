package server

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (s *Server) handleGallery(c *fiber.Ctx) error {
    ids, err := s.idx.FilterFileCid(50)
    if err != nil {
        zap.S().Errorf("filter file cid err: %s", err)
        return err
    }

    return c.JSON(ids)
}