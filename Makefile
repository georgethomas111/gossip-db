build:
	go build cmd/gossip/gossip.go

compose:
	docker compose up --force-recreate --build
