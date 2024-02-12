package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/hexisa_go_nal_todo/internal/common/config"
	"github.com/hexisa_go_nal_todo/internal/common/server"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg := config.NewConfig()

	mux, cleanup, err := server.NewMux(ctx, cfg)
	defer cleanup()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.AppPort))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Server.AppPort, err)
	}

	s := server.NewServer(l, mux)
	return s.Run(ctx)
}
