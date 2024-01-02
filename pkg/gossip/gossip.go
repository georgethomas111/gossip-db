package gossip

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/georgethomas111/gossip-db/pkg/node"
	"github.com/georgethomas111/gossip-db/pkg/stats"
)

var HeartBeatMs = 500

type Client interface {
	List() (map[string]*node.Row, error)
}

func Talk(n *node.Node, client Client) error {
	listedMap, err := client.List()
	if err != nil {
		return errors.New("Gossip client List() failed " + err.Error())
	}

	for key, row := range listedMap {
		newRow := node.MergeRows(n.Get(key), row)
		n.Put(key, newRow)
	}

	stats.RowsInDatabase.Set(float64(len(listedMap)))

	return nil
}

type Gossip struct {
	Others    []string
	Instance  *node.Node
	HeartBeat *time.Ticker
	clients   []*JSONClient
}

func New(instance *node.Node, others []string) *Gossip {
	var t []*JSONClient

	stats.NodeCount.Set(float64(len(others)))
	for _, addr := range others {
		t = append(t, NewJSONClient(addr))
	}

	g := &Gossip{
		Others:    others,
		Instance:  instance,
		HeartBeat: time.NewTicker(time.Duration(HeartBeatMs) * time.Millisecond),
		clients:   t,
	}

	// context so that graceful shutdowns can be added some day.
	g.runGossips(context.Background())

	return g

}

func (g *Gossip) runGossips(ctx context.Context) {
	for {
		select {
		case <-g.HeartBeat.C:
			for _, c := range g.clients {
				err := Talk(g.Instance, c)
				if err != nil {
					fmt.Println("Error talking to client " + err.Error())
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
