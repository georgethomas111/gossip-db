package gossip

import (
	"encoding/json"
	"testing"

	"github.com/georgethomas111/gossip-db/pkg/node"
)

type MockClient struct {
	ServerNode *node.Node
}

func (m *MockClient) List() (map[string]*node.Row, error) {
	JSONData, err := m.ServerNode.ListJSON()
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

func TestGossip(t *testing.T) {
	node1, err := node.New()
	if err != nil {
		t.Error("Error creating node ", err.Error())
	}
	node1.PutVal("1", []byte("a"))
	node1.PutVal("3", []byte("c"))

	node2, err := node.New()
	if err != nil {
		t.Error("Error creating node ", err.Error())
		return
	}
	node2.PutVal("2", []byte("b"))
	client := &MockClient{
		ServerNode: node2,
	}

	err = Gossip(node1, client)
	if err != nil {
		t.Error("Error Gossiping ", err.Error())
		return
	}

	row := node1.Get("2")
	if row == nil {
		t.Error("Got nil row when a row is expected")
		return
	}

	if string(row.Value) != "b" {
		t.Errorf("Expected b, Got %v \n", string(row.Value))
	}
}
