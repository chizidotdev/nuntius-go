build:
	go build -o build/bin/main ./cmd/main.go

postgres-cli:
	docker exec -it nuntius-db psql --username=root --dbname=nuntius

templ:
	templ fmt . && templ generate

server: templ
	air

.PHONY: build postgres-cli server templ
