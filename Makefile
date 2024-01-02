build:
	GOOS=wasip1 GOARCH=wasm go build -o gossip.wasm cmd/gossip/gossip.go

wasm-opt:
	wasm-opt -O --enable-bulk-memory gossip.wasm -o gossip-opt.wasm
