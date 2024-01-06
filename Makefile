build:
	go build -o bin/copia ./main.go

postgres:
	docker run --name nuntius -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

postgres-cli:
	docker exec -it nuntius psql --username=root --dbname=nuntius

createdb:
	docker exec -it nuntius createdb --username=root --owner=root nuntius

dropdb:
	docker exec -it nuntius dropdb nuntius

server:
	make templ && air

templ:
	templ fmt . && templ generate

.PHONY: build postgres postgres-cli createdb dropdb server templ