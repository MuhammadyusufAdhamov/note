DB_URL=postgresql://postgres:7@localhost:5432/note?sslmode=disable

swag-init:
	swag init -g api/api.go -o api/docs

start:
	go run main.go

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

.PHONY: start migrateup migratedown