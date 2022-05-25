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

include .env
export
test:	
	go test ./...