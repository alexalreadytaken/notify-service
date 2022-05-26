help:
	$(info TODO)

db:
	docker-compose up -d db

 up-rebuild:
	docker-compose build --force-rm --no-cache notifyer
	docker-compose up -d

upall:
	docker-compose up -d

stopall:
	docker-compose stop

logs:
	docker-compose logs

swag-gen:
	swag init -g cmd/main.go

include .env
export
notifyer:
	go run cmd/main.go

test:	
	go test ./...