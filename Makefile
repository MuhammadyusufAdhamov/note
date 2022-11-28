DB_URL=postgresql://postgres:7@localhost:5432/note?sslmode=disable

start:
	go run main.go

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

.PHONY: start migrateup migratedown