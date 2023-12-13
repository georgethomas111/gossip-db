package node

import "sync"

type Node struct {
	data map[string][]byte
	m    sync.Mutex
}

func (n *Node) Get(key string) []byte {
	n.m.Lock()
	defer n.m.Unlock()
	return n.data[key]
}

func (n *Node) Put(key string, val []byte) {
	n.m.Lock()
	defer n.m.Unlock()
	n.data[key] = val
}

func (n *Node) ListKeys() []string {
	n.m.Lock()
	defer n.m.Unlock()
	var keys []string
	for key, _ := range n.data {
		keys = append(keys, key)
	}
	return keys
}

func (n *Node) ListVals() [][]byte {
	n.m.Lock()
	defer n.m.Unlock()
	var vals [][]byte
	for _, val := range n.data {
		vals = append(vals, val)
	}
	return vals

}

func New() (*Node, error) {
	var lock sync.Mutex
	return &Node{
		data: make(map[string][]byte),
		m:    lock,
	}, nil

}
