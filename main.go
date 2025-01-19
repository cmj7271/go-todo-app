package main

import (
	"context"
	"fmt"
	"github.com/cmj7271/go-todo-app/config"
	"log"
	"net"
	"os/signal"
	"syscall"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

func run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.New()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
		return err
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v\n", url)

	mux := newMux()
	s := NewServer(l, mux)
	return s.Run(ctx)
}
