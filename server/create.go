package server

import (
	"io"
	"mime"
	"net/textproto"

	"go.uber.org/zap"

	"github.com/mayocream/pastebin-ipfs/pkg/index"
	"github.com/mayocream/pastebin-ipfs/pkg/ipfs"
)

type file struct {
	Name     string
	MIMEType string
	Reader   io.Reader
	Size     int64
}

func (s *Server) metadata(files ...*file) (objs []Object) {
	for _, f := range files {
		obj := Object{
			Name:     f.Name,
			MIMEType: f.MIMEType,
			Size:     f.Size,
		}
		objs = append(objs, obj)
	}
	return
}

func (s *Server) creates(files ...*file) (cid string, err error) {
	fs := make([]*ipfs.File, 0, len(files))
	for _, f := range files {
		fs = append(fs, &ipfs.File{
			Name:   f.Name,
			Reader: f.Reader,
		})
	}

	res, err := s.ipc.Add(fs...)
	if err != nil {
		return
	}

	for _, obj := range res.Objects {
		ot := index.ObjectTypeFile
		if obj.Name == "" {
			ot = index.ObjectTypeDir
		} else if obj.Name == metadataFileName {
			ot = index.ObjectTypeMeta
		}
		if err = s.idx.SetExist(obj.Hash, ot); err != nil {
			zap.S().With("cid", obj.Hash).Errorf("idx set err: %s", err)
			return
		}
	}

	return
}

func mediaTypeOrDefault(header textproto.MIMEHeader) string {
	mediaType, _, err := mime.ParseMediaType(header.Get("Content-Type"))
	if err != nil {
		return "application/octet-stream"
	}
	return mediaType
}
