package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.miganbox.com/migan/surl/backend"
	"git.miganbox.com/migan/surl/repository"
)

var (
	Version   = "0.0.0"
	Branch    = "local"
	Commit    = "000000"
	UpdatedAt = "000000" // +%y%m%d
)

func main() {
	srv := backend.New()

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	slog.Info("surl started!",
		"version", fmt.Sprintf("%s-backend", Version),
		"branch", Branch,
		"commit", Commit,
		"updated_at", UpdatedAt,
	)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		// TODO: Use slog
		fmt.Println(err)
	}

	if err := repository.GetDatabase().Close(); err != nil {
		// TODO: Use slog
		fmt.Println(err)
	}
}
