package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/georgethomas111/gossip-db/pkg/node"
	"github.com/georgethomas111/gossip-db/pkg/swim"
	"github.com/gorilla/mux"
)

const DefaultPortVar = 8000
const swimPortDiff = 1000

var portVar = DefaultPortVar
var instance *node.Node
var known = "node-known:7000"

func init() {
	flag.IntVar(&portVar, "port", portVar, "The port the web browser will be looking for. Swim port will be 1000 less than this.")
	flag.BoolVar(&swim.FirstNode, "firstNode", swim.FirstNode, "Boolean indicating if this is the known node.")
	flag.StringVar(&known, "known", known, "The Address of the known node. Swim port will be 1000 less than this.")
	flag.Parse()
}

func routeMap() map[string]func(http.ResponseWriter, *http.Request) {
	routes := instance.Routes()
	//	routes["/members"] = Members
	//	routes["/addMembers/{member}"] = AddMembers
	return routes
}

func routeHandler(routes map[string]func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for path, _ := range routes {
			fmt.Fprintln(w, path)
		}
	}
}

/*
func otherAddresses(otherNodes []string) []string {
	var others []string
	for _, addr := range otherNodes {
		nodePort := swimPort + swimPortDiff
		if nodePort != portVar {
			otherAddr := "0.0.0.0" + fmt.Sprintf(":%d", nodePort)
			others = append(others, otherAddr)
		}
	}
	return others
}
*/

func main() {
	// Let swim port be always SwimPortDiff less to start with.
	swimPort := portVar - swimPortDiff

	var err error
	instance, err = node.New()
	if err != nil {
		panic("Initializing node " + err.Error())
	}

	r := mux.NewRouter()
	routes := routeMap()
	r.HandleFunc("/", routeHandler(routes))

	for path, fn := range routes {
		r.HandleFunc(path, fn)
	}

	otherNodes, err := swim.New(swimPort, known)
	if err != nil {
		panic("Swim could not be started " + err.Error())
	}

	srv := &http.Server{
		Addr:    "0.0.0.0:" + fmt.Sprintf("%d", portVar),
		Handler: r,
	}
	srvAddr := srv.Addr
	fmt.Println("Serving routes over ", srvAddr)
	fmt.Println("Swim port at  ", swimPort)

	//	others := otherAddresses(otherNodes)
	fmt.Println("Address of nodes = ", otherNodes)
	//	go gossip.New(instance, others)

	srv.ListenAndServe()
}
