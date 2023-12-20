package gossip

import (
	"errors"

	"github.com/georgethomas111/gossip-db/pkg/node"
)

type Client interface {
	List() (map[string]*node.Row, error)
}

func Gossip(n *node.Node, client Client) error {
	listedMap, err := client.List()
	if err != nil {
		return errors.New("Gossip client List() failed " + err.Error())
	}

	for key, row := range listedMap {
		newRow := node.MergeRows(n.Get(key), row)
		n.Put(key, newRow)
	}

	return nil
}

/*
import (
	"context"
	"time"
)

var HeartBeatMs = 500

type Transport interface {
}

//    n1:8080   n2
//    findOthers

// ------------------------------------------------>
//                   t ->

type distributed struct {
	node *Node
}

func Distributed(dnsNodeAddr string) *distributed {
	return &distributed{
		dnsNodeAddr:    dnsNodeAddr,
		heartBeatTimer: time.NewTimer(HeartBeatMs * time.Milliseconds),
	}
}

// Each node should gossip periodically with a heart beat
func (d *distributed) Join(ctx context.Context, instance *Node) *Gossip {
}

func (g *Gossip) AddNode(node string) {
	g.nodeAddresses = append(g.nodeAddresses, node)
}

func (g *Gossip) runGossips(ctx context.Context) {
	for {
		select {
		case <-g.heartBeatTimer.C:
			// iterate throught the nodeAddresses
			// Call the other nodes and fill up the current nodes data
			return
		case <-ctx.Done():
			return
		}
	}
}
*/
