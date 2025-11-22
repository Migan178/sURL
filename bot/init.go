package main

import (
	_ "github.com/Migan178/surl-bot/commands"
	_ "github.com/Migan178/surl-bot/components"
	"github.com/Migan178/surl-bot/configs"
	"github.com/Migan178/surl-bot/handler"
	_ "github.com/Migan178/surl-bot/modals"
	"github.com/bwmarrin/discordgo"
)

var dg *discordgo.Session

func init() {
	dg, _ = discordgo.New("Bot " + configs.GetConfig().Bot.Token)

	// Handler
	go dg.AddHandler(handler.InteractionCreate)
}
