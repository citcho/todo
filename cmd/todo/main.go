package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/hexisa_go_nal_todo/internal/pkg/config"
	"github.com/hexisa_go_nal_todo/internal/pkg/database"
	"github.com/hexisa_go_nal_todo/internal/pkg/server"
)

func main() {
	log.Println("starting server...")
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg := config.NewConfig()
	closeDB := database.NewDB(ctx, cfg.DB)
	defer closeDB()
	mux := server.NewMux(ctx, cfg)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.AppPort))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Server.AppPort, err)
	}

	s := server.NewServer(l, mux)
	return s.Run(ctx)
}
