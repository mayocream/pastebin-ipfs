package server

import (
	"io"
	"mime"
	"net/textproto"

	"github.com/spf13/cast"
	"go.uber.org/zap"

	"github.com/mayocream/pastebin-ipfs/pkg/ipfs"
)

type file struct {
	Name     string
	MIMEType string
	Reader   io.Reader
}

func (s *Server) creates(files ...*file) ([]Object, error) {
	fs := make([]*ipfs.File, 0, len(files))
	mimetypes := make(map[string]string)
	for _, f := range files {
		mimetypes[f.Name] = f.MIMEType
		fs = append(fs, &ipfs.File{
			Name:   f.Name,
			Reader: f.Reader,
		})
	}

	res, err := s.ipc.Add(fs...)
	if err != nil {
		return nil, err
	}

	for _, obj := range res.Objects {
		if err := s.idx.SetExist(obj.Hash); err != nil {
			zap.S().With("cid", obj.Hash).Errorf("idx set err: %s", err)
			return nil, err
		}
	}

	objs := make([]Object, 0, len(res.Objects))
	for _, v := range res.Objects {
		obj := Object{
			Cid:      v.Hash,
			Name:     v.Name,
			MIMEType: mimetypes[v.Name],
			Size:     cast.ToInt64(v.Size),
		}
		objs = append(objs, obj)
	}
	return objs, nil
}

func mediaTypeOrDefault(header textproto.MIMEHeader) string {
	mediaType, _, err := mime.ParseMediaType(header.Get("Content-Type"))
	if err != nil {
		return "application/octet-stream"
	}
	return mediaType
}
