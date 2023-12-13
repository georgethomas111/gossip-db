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

func init() {
	flag.StringVar(&portVar, "port", portVar, "The port the web browser will be looking for")
	flag.Parse()
}

func PutKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	val := vars["value"]
	fmt.Fprintln(w, "key = ", key)
	fmt.Fprintln(w, "val = ", val)
	instance.Put(key, []byte(val))

}

func GetKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	fmt.Fprintln(w, "key = ", key)
	fmt.Fprintln(w, "Value =", string(instance.Get(key)))
}

func ListVals(w http.ResponseWriter, r *http.Request) {
	respVals := instance.ListVals()
	for _, val := range respVals {
		fmt.Fprintln(w, string(val))
	}
}

func ListKeys(w http.ResponseWriter, r *http.Request) {
	respKeys := instance.ListKeys()
	for _, key := range respKeys {
		fmt.Fprintln(w, key)
	}
}

func routeMap() map[string]func(http.ResponseWriter, *http.Request) {
	routes := make(map[string]func(http.ResponseWriter, *http.Request))

	routes["/put/{key}/{value}"] = PutKey
	routes["/get/{key}"] = GetKey
	routes["/listKeys"] = ListKeys
	routes["/listVals"] = ListVals
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

	fmt.Println("Serving routes over ", portVar)
	srv := &http.Server{
		Addr:    "0.0.0.0:" + portVar,
		Handler: r,
	}

	srv.ListenAndServe()

}
