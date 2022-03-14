bup:
	docker-compose build
	docker-compose up

bld:
	docker-compose build

up:
	docker-compose up

run:
	swag init -g main.go --output docs/wtc
	go run main.go
