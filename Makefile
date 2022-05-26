help:
	$(info TODO)

db:
	docker-compose up -d db

full:
	docker-compose up -d

stopall:
	docker-compose stop

logs:
	docker-compose logs

swag-gen:
	swag init -g cmd/main.go

include .env
export
test:	
	go test ./...