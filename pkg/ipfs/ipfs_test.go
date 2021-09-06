package ipfs

import (
	"sync"
	"testing"
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

	cid, err := cli.Add([]byte(`
	# Snale cute
	`))

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("cid: %s", cid)
}

func TestClient_Cat(t *testing.T) {
	cli := getTestClient()

	file, err := cli.Cat("QmWLr7ca8U8CWiNZEZjd22PY87i85C8tioPPKa8gcnVcaj")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("content: %s", file)
}
