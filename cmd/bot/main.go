package main

import (
	"os"
	"os/signal"
	"syscall"

	"git.miganbox.com/migan/surl/bot"
)

func main() {
	b := bot.New()

	if err := b.Open(); err != nil {
		panic(err)
	}

	defer b.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
