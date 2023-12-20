package node

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (n *Node) PutKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	val := vars["value"]
	fmt.Fprintln(w, "key = ", key)
	fmt.Fprintln(w, "val = ", val)
	n.PutVal(key, []byte(val))

}

func (n *Node) GetKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	fmt.Fprintln(w, "key = ", key)
	row := n.Get(key)
	if row == nil {
		fmt.Fprintln(w, "Row with key specified not found")
		return
	}

	fmt.Fprintln(w, "Value = ", string(row.Value))
	fmt.Fprintln(w, "PutTimestamp = ", row.PutTimestamp)

}

func (n *Node) ListVals(w http.ResponseWriter, r *http.Request) {
	respVals := n.listVals()
	for _, val := range respVals {
		fmt.Fprintln(w, string(val))
	}
}

func (n *Node) ListKeys(w http.ResponseWriter, r *http.Request) {
	respKeys := n.listKeys()
	for _, key := range respKeys {
		fmt.Fprintln(w, key)
	}
}

func (n *Node) Routes() map[string]func(http.ResponseWriter, *http.Request) {
	routes := make(map[string]func(http.ResponseWriter, *http.Request))
	routes["/put/{key}/{value}"] = n.PutKey
	routes["/get/{key}"] = n.GetKey
	routes["/listKeys"] = n.ListKeys
	routes["/listVals"] = n.ListVals
	return routes
}
