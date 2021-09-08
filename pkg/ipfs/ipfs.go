package ipfs

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"

	shell "github.com/ipfs/go-ipfs-api"
	files "github.com/ipfs/go-ipfs-files"
)

// File ipfs file to upload
type File struct {
	Reader io.Reader
	Name   string
}

// Client IPFS client
type Client struct {
	host string
	sh   *shell.Shell
}

// NewClient create new IPFS client
func NewClient(host string) (*Client, error) {
	sh := shell.NewShell("localhost:5001")
	c := &Client{
		host: host,
		sh:   sh,
	}
	return c, nil
}

// Ping test network
func (c *Client) Ping() error {
	if !c.sh.IsUp() {
		return errors.New("server down")
	}
	return nil
}

// Object ...
type Object struct {
	Bytes string
	Hash  string
	Name  string
	Size  string
}

// AddResult ...
type AddResult struct {
	Cid     string   `json:"cid"`
	Objects []Object `json:"Objects"`
}

// Add add files
func (c *Client) Add(srcs ...*File) (*AddResult, error) {
	nodes := make(map[string]files.Node, len(srcs))
	for _, src := range srcs {
		nodes[src.Name] = files.NewReaderFile(src.Reader)
	}

	sf := files.NewMapDirectory(nodes)
	slf := files.NewSliceDirectory([]files.DirEntry{files.FileEntry("", sf)})
	reader := files.NewMultiFileReader(slf, true)

	resp, err := c.sh.Request("add").
		Option("recursive", true).
		Body(reader).
		Send(context.Background())
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	result := new(AddResult)
	dec := json.NewDecoder(resp.Output)
	for {
		var obj Object
		err = dec.Decode(&obj)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		result.Objects = append(result.Objects, obj)
		result.Cid = obj.Hash
	}

	return result, nil
}

// Cat cat file
func (c *Client) Cat(cid string) ([]byte, error) {
	src, err := c.sh.Cat(cid)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, src)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), err
}

// Cat cat file
func (c *Client) CatStream(cid string) (io.ReadCloser, error) {
	src, err := c.sh.Cat(cid)
	if err != nil {
		return nil, err
	}

	return src, err
}
