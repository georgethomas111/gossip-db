# Steps to Run

```
// Expose the current directory over the browser.
$ go run cmd/gossip/gossip.go

```

On your browser go to http://localhost:8080/ to see the options and data that is present.

# README

* First create a simple map with put and get with endpoints that can achieve it. 
* The endpoints listen on ports and has two versions of it.
* Add gossip package with handcoded list of members.

Gossiping works

# Proof it works.
Run one node at 8081 gossiping with 8080

```
george@pop-os:~/workspace/gossip-db$ go run cmd/gossip/gossip.go -port 8081 -other 8080
Serving routes over  0.0.0.0:8081
```

Run other node at 8080 gossiping with 8081
```
george@pop-os:~/workspace/gossip-db$ go run cmd/gossip/gossip.go -port 8080 -other 8081
Serving routes over  0.0.0.0:8080
```

```
george@pop-os:~/workspace/gossip-db$ curl -X GET http://localhost:8080/
/put/{key}/{value}
/get/{key}
/listKeys
/listVals
/listJSON
george@pop-os:~/workspace/gossip-db$ curl -X GET http://localhost:8080/put/key1/787
key =  key1
val =  787
george@pop-os:~/workspace/gossip-db$ curl -X GET http://localhost:8081/get/key1
key =  key1
Value =  787
PutTimestamp =  2023-12-20 20:18:55.201666149 -0800 PST
george@pop-os:~/workspace/gossip-db$ curl -X GET http://localhost:8080/get/key1
key =  key1
Value =  787
PutTimestamp =  2023-12-20 20:18:55.201666149 -0800 PST
george@pop-os:~/workspace/gossip-db$ 
```
