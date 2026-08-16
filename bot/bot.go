package bot

import (
	"git.miganbox.com/migan/surl/bot/commands"
	_ "git.miganbox.com/migan/surl/bot/components"
	"git.miganbox.com/migan/surl/bot/handler"
	_ "git.miganbox.com/migan/surl/bot/modals"
	"git.miganbox.com/migan/surl/configs"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	s                 *discordgo.Session
	globalCmds        []*discordgo.ApplicationCommand
	developerOnlyCmds []*discordgo.ApplicationCommand
}

func New() *Bot {
	s, _ := discordgo.New("Bot " + configs.GetConfig().Bot.Token)

	// Handler
	go s.AddHandler(handler.InteractionCreate)

	var globalCmds []*discordgo.ApplicationCommand
	var developerOnlyCmds []*discordgo.ApplicationCommand
	for _, cmd := range commands.GetDiscommand().Commands {
		if cmd.Flags&commands.CommandFlagsIsDeveloperOnlyCommand != 0 {
			developerOnlyCmds = append(developerOnlyCmds, cmd.ApplicationCommand)
			continue
		}

		globalCmds = append(globalCmds, cmd.ApplicationCommand)
	}

	return &Bot{s, globalCmds, developerOnlyCmds}
}

func (b *Bot) Open() error {
	if err := b.s.Open(); err != nil {
		return err
	}

	_, err := b.s.ApplicationCommandBulkOverwrite(b.s.State.User.ID, "", b.globalCmds)
	if err != nil {
		return err
	}

	developerOnlyCommandGuildID := configs.GetConfig().Bot.DeveloperOnlyCommandGuildID
	if len(b.developerOnlyCmds) != 0 && developerOnlyCommandGuildID != "" {
		_, err = b.s.ApplicationCommandBulkOverwrite(b.s.State.User.ID, developerOnlyCommandGuildID, b.developerOnlyCmds)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Bot) Close() error {
	return b.s.Close()
}
