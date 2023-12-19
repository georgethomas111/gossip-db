package node

import (
	"sync"
)

type Node struct {
	data map[string]*Row
	m    sync.Mutex
}

func (n *Node) Get(key string) *Row {
	n.m.Lock()
	defer n.m.Unlock()
	return n.data[key]
}

func (n *Node) Put(key string, val []byte) {
	d := NewRow(val)
	n.m.Lock()
	defer n.m.Unlock()
	n.data[key] = d
}

func (n *Node) listKeys() []string {
	n.m.Lock()
	defer n.m.Unlock()
	var keys []string
	for key, _ := range n.data {
		keys = append(keys, key)
	}
	return keys
}

func (n *Node) listVals() [][]byte {
	n.m.Lock()
	defer n.m.Unlock()
	var vals [][]byte
	for _, row := range n.data {
		vals = append(vals, row.Value)
	}
	return vals
}

func New() (*Node, error) {
	var lock sync.Mutex
	return &Node{
		data: make(map[string]*Row),
		m:    lock,
	}, nil
}
