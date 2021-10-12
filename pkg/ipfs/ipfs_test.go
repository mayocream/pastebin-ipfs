package ipfs

import (
	"strings"
	"sync"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

var testOnce sync.Once

var testClient *Client

func getTestClient() *Client {
	testOnce.Do(func() {
		var err error
		testClient, err = NewClient("127.0.0.1:5001")
		if err != nil {
			panic(err)
		}
	})
	return testClient
}

func TestClient_Ping(t *testing.T) {
	cli := getTestClient()

	err := cli.Ping()
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Add(t *testing.T) {
	cli := getTestClient()

	cid, err := cli.Add(&File{
		Name:   "song.txt",
		Reader: strings.NewReader("唱不完一首歌"),
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("cid: %s", spew.Sdump(cid))
}

func TestClient_Cat(t *testing.T) {
	cli := getTestClient()

	content, err := cli.Cat("QmSpVsNZaamFRh67SMWPqaGrCHetshAhfn7a16it94JQw9/song.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("content: %s", content)
}
