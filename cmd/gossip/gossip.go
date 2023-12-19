package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/georgethomas111/gossip-db/pkg/node"
	"github.com/gorilla/mux"
)

var portVar = "8080"
var instance *node.Node
var dnsNodeAddr = "127.0.0.1:8080"

func init() {
	flag.StringVar(&portVar, "port", portVar, "The port the web browser will be looking for")
	flag.StringVar(&dnsNodeAddr, "dnsNodeAddr", dnsNodeAddr, "Its the node to call when you just started")
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

	srv.ListenAndServe()
}
