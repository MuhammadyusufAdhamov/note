package main

import (
	"fmt"
	"log"

	"github.com/MuhammadyusufAdhamov/note/config"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)


func main() {
	cfg := config.Load(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	_, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	log.Print("Server stopped")
}