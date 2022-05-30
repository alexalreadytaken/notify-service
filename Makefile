help:
	$(info commands:)
	$(info )
	$(info -make db = запускает контейнер пг в фоне)
	$(info )
	$(info -make up-rebuild = ребилдит контейнер с бэком)
	$(info )
	$(info -make upall = поднять все контейнеры в фоне)
	$(info )
	$(info -make stopall = остановить все контейнеры)
	$(info )
	$(info -make logs = выводит логи контейнеров)
	$(info )
	$(info -make swag-gen = обновляет спецификацию swagger)
	$(info )
	$(info -make notifyer = запускает бэк без контейнера)
	$(info )
	$(info -make test = запускает тесты бэка)	
	$(info )

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