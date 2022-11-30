package main

import (
	"fmt"
	"log"

	"github.com/MuhammadyusufAdhamov/note/api"
	"github.com/MuhammadyusufAdhamov/note/config"
	"github.com/MuhammadyusufAdhamov/note/storage"
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

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	strg := storage.NewStoragePg(psqlConn)

	apiServer := api.New(&api.RouterOptions{
		Cfg: &cfg,
		Storage: strg,
	})
	err = apiServer.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("failed to tun server error %v", err)
	}
	
	log.Print("Server stopped")
}