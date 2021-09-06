package ipfs

import (
	"bytes"
	"errors"
	"io"

	shell "github.com/ipfs/go-ipfs-api"
)

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

// Add add bytes
func (c *Client) Add(file []byte) (string, error) {
	cid, err := c.sh.Add(bytes.NewReader(file))
	return cid, err
}

// Cat cat file
func (c *Client) Cat(cid string) ([]byte, error) {
	src, err := c.sh.Cat(cid)
	if err != nil {
		return nil, err
	}

	// TODO max read size
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