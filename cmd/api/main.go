package main

import (
	"fmt"
	"log"

	"github.com/mostafizurrahaman/go_server/internal/config"
	"github.com/mostafizurrahaman/go_server/internal/db"
	"github.com/mostafizurrahaman/go_server/internal/server"
)

func main() {

	// ?? Load environment variables
	cfg, err := config.Load()

	if err != nil {
		log.Println("Failed to load env. %w", err)
	}

	// ?? Setup database connection:
	mgClient, database, err := db.Connect(cfg)
	if err != nil {
		fmt.Println("Failed to connect db. %w", err)
	}

	r := server.NewRouter(database)

	addr := fmt.Sprintf(":%s", cfg.Port)
	if err := r.Run(addr); err != nil {
		log.Println("Failed to run server. %w", err)
	}

	defer db.Disconnect(mgClient)

}
