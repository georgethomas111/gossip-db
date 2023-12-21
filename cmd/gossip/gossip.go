package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/georgethomas111/gossip-db/pkg/gossip"
	"github.com/georgethomas111/gossip-db/pkg/node"
	"github.com/gorilla/mux"
)

var portVar = "8080"
var otherPort = "8081"
var instance *node.Node

func init() {
	flag.StringVar(&portVar, "port", portVar, "The port the web browser will be looking for")
	flag.StringVar(&otherPort, "other", otherPort, "The port of the other port that is listening")
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

func main() {
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

	srv := &http.Server{
		Addr:    "0.0.0.0:" + portVar,
		Handler: r,
	}
	srvAddr := srv.Addr
	fmt.Println("Serving routes over ", srvAddr)

	otherAddr := "0.0.0.0:" + otherPort

	var others []string = []string{otherAddr}
	go gossip.New(instance, others)

	srv.ListenAndServe()
}
