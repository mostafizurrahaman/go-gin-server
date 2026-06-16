package main

import (
	"fmt"
	"log"

	"github.com/mostafizurrahaman/go_server/internal/server"
)

func main() {


	



	r := server.NewRouter()

	addr := fmt.Sprint(":5000")
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to run server. %w", err)
	}

}
