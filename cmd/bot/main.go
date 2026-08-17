package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"git.miganbox.com/migan/surl/bot"
)

var (
	Version   = "0.0.0"
	Branch    = "local"
	Commit    = "000000"
	UpdatedAt = "000000" // +%y%m%d
)

func main() {
	b := bot.New()

	if err := b.Open(); err != nil {
		panic(err)
	}

	slog.Info("surl started!",
		"version", fmt.Sprintf("%s-bot", Version),
		"branch", Branch,
		"commit", Commit,
		"updated_at", UpdatedAt,
	)

	defer b.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
