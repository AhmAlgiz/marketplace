build:
	docker-compose build marketplace

run:
	docker-compose up marketplace

migrate:
	migrate -path ./schema -database 'postgres://postgres:123456@0.0.0.0:5432/postgres?sslmode=disable' up

