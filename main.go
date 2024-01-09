package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"dalle/api"
	"dalle/config"
	"dalle/extapi/openai"
	"dalle/repository"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

const idleTimeout = 5 * time.Second

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	openaiService := openai.New(cfg.OpenAI.APIKey)
	repo, err := repository.New(cfg.Postgres)
	if err != nil {
		panic(err)
	}

	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	api.ApplyRoutes(app, openaiService, repo)

	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
			log.Fatalf("cant start server: %v", err)
		}
	}()

	<-ctx.Done()

	if err := app.Shutdown(); err != nil {
		log.Fatalf("cant stop server: %v", err)
	}

	log.Info("server stopped")
}
