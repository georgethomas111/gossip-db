package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/georgethomas111/gossip-db/pkg/node"
)

var portVar = "8080"
var n *node.Node

func init() {
	flag.StringVar(&portVar, "port", portVar, "The port the web browser will be looking for")
	flag.Parse()
}

//Displays a paticular file in a folder
func DispFile(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(r.URL.Path[1:])
	if err == nil {
		fmt.Fprintf(w, "%s", data)
	} else {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// Lists the Files in the current folder.
func ListFiles(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(".")
	if err == nil {
		for _, file := range files {
			fmt.Fprintf(w, "<a href='%s'>%s</a><br/>", file.Name(), file.Name())
		}
	} else {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func Members(w http.ResponseWriter, r *http.Request) {
	for _, member := range n.Members() {
		fmt.Fprintf(w, "Member: %s %s\n", member.Name, member.Addr)
	}

}

func main() {
	var err error
	n, err = node.New()
	if err != nil {
		panic("Initializing node " + err.Error())
	}

	http.HandleFunc("/members", Members)
	http.HandleFunc("/list", ListFiles)
	http.HandleFunc("/", DispFile)

	fmt.Println("Serving files in this director over port ", portVar)
	http.ListenAndServe(":"+portVar, nil)
}
