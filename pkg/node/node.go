package node

import "sync"

type Node struct {
	data map[string][]byte
	m    sync.Mutex
}

func (n *Node) Get(key string) {
	n.m.Lock()
	defer n.m.Unlock()
	data[key] = val
}

func (n *Node) Put(key string, val []byte) {
	n.m.Lock()
	defer n.m.Unlock()
	data[key] = val
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
	return &Node{
		data: make(map[string][]byte),
	}, nil

}
