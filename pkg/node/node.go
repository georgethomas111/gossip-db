package node

import (
	"encoding/json"
	"sync"
)

type Node struct {
	Data map[string]*Row
	m    sync.Mutex
}

func (n *Node) Get(key string) *Row {
	n.m.Lock()
	defer n.m.Unlock()
	return n.Data[key]
}

func (n *Node) Put(key string, row *Row) {
	n.m.Lock()
	defer n.m.Unlock()
	n.Data[key] = row
}

func (n *Node) PutVal(key string, val []byte) {
	d := NewRow(val)
	n.m.Lock()
	defer n.m.Unlock()
	n.Data[key] = d
}

func (n *Node) listKeys() []string {
	n.m.Lock()
	defer n.m.Unlock()
	var keys []string
	for key, _ := range n.Data {
		keys = append(keys, key)
	}
	return keys
}

func (n *Node) listVals() [][]byte {
	n.m.Lock()
	defer n.m.Unlock()
	var vals [][]byte
	for _, row := range n.Data {
		vals = append(vals, row.Value)
	}
	return vals
}

func (n *Node) ListJSON() ([]byte, error) {
	return json.Marshal(n.Data)
}

func New() (*Node, error) {
	var lock sync.Mutex
	return &Node{
		Data: make(map[string]*Row),
		m:    lock,
	}, nil
}
