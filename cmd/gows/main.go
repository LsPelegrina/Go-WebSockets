package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/LsPelegrina/Go-WebSockets/internal/api"
	"github.com/LsPelegrina/Go-WebSockets/internal/store/pgstore"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, fmt.Sprintf(
		"port=%s user=%s password=%s dbname=%s host=%s",
		os.Getenv("WS_DATABASE_PORT"),
		os.Getenv("WS_DATABASE_USER"),
		os.Getenv("WS_DATABASE_PASSWORD"),
		os.Getenv("WS_DATABASE_NAME"),
		os.Getenv("WS_DATABASE_HOST"),
	))
	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	handler := api.NewHandler(pgstore.New(pool))

	go func() {
		if err := http.ListenAndServe(":8000", handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
