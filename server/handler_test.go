package server

import (
	"log"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/pretty"

	"github.com/mayocream/pastebin-ipfs/pkg/index"
	"github.com/mayocream/pastebin-ipfs/pkg/ipfs"
)

var once sync.Once

func testServer() {
	once.Do(func() {
		idx, err := index.NewIndex("/tmp/pastebin_ipfs")
		if err != nil {
			log.Fatal(err)
		}
		ipc, err := ipfs.NewClient("localhost:5001")
		if err != nil {
			log.Fatal(err)
		}
		srv := New(&Config{
			IPFSClient: ipc,
			Index:      idx,
		})
		go srv.Start(":3940")
		time.Sleep(1 * time.Second)
	})
}

func TestServer_handleUpload(t *testing.T) {
	testServer()

	resp, err := resty.New().R().
		SetFileReader("file1", "plain.txt", strings.NewReader("hello world!")).
		SetFileReader("file2", "wow.txt", strings.NewReader("wow!")).
		SetFormData(map[string]string{
			"author": "tester",
		}).Post("http://127.0.0.1:3940/api/v1/upload")

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("resp: %s", pretty.Pretty(resp.Body()))
}
