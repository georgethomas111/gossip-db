package node

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PutRoute      = "/put/{key}/{value}"
	GetRoute      = "/get/{key}"
	ListKeysRoute = "/listKeys"
	ListValsRoute = "/listVals"
	ListJSONRoute = "/listJSON"
)

func (n *Node) PutKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	val := vars["value"]
	fmt.Fprintln(w, "key = ", key)
	fmt.Fprintln(w, "val = ", val)
	n.PutVal(key, []byte(val))

}

func (n *Node) GetKeyHandler(w http.ResponseWriter, r *http.Request) {
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

func (n *Node) ListValsHandler(w http.ResponseWriter, r *http.Request) {
	respVals := n.listVals()
	for _, val := range respVals {
		fmt.Fprintln(w, string(val))
	}
}

func (n *Node) ListKeysHandler(w http.ResponseWriter, r *http.Request) {
	respKeys := n.listKeys()
	for _, key := range respKeys {
		fmt.Fprintln(w, key)
	}
}

func (n *Node) ListJSONHandler(w http.ResponseWriter, r *http.Request) {
	j, err := n.ListJSON()
	if err != nil {
		fmt.Fprintln(w, "Error Listing "+err.Error())
	}
	fmt.Fprintln(w, string(j))
}

func (n *Node) Routes() map[string]func(http.ResponseWriter, *http.Request) {
	routes := make(map[string]func(http.ResponseWriter, *http.Request))
	routes[PutRoute] = n.PutKeyHandler
	routes[GetRoute] = n.GetKeyHandler
	routes[ListKeysRoute] = n.ListKeysHandler
	routes[ListValsRoute] = n.ListValsHandler
	routes[ListJSONRoute] = n.ListJSONHandler

	return routes
}
