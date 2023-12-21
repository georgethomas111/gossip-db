package gossip

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/georgethomas111/gossip-db/pkg/node"
)

type JSONClient struct {
	Addr string
}

func (j *JSONClient) callRemote() ([]byte, error) {
	endpoint := "http://" + j.Addr + node.ListJSONRoute
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func (j *JSONClient) List() (map[string]*node.Row, error) {
	JSONData, err := j.callRemote()
	if err != nil {
		return nil, err
	}

	data := make(map[string]*node.Row)
	err = json.Unmarshal(JSONData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewJSONClient(addr string) *JSONClient {
	return &JSONClient{
		Addr: addr,
	}
}
