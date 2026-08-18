package bot

import (
	"context"

	_ "git.miganbox.com/migan/surl/bot/commands"
	"git.miganbox.com/migan/surl/bot/loader"
	"git.miganbox.com/migan/surl/configs"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
)

type Bot struct {
	s *bot.Client
}

func New() *Bot {
	s, _ := disgo.New(configs.GetConfig().Bot.Token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuilds,
			),
		),
		bot.WithEventListeners(loader.Loader().Router()),
	)

	return &Bot{s}
}

func (b *Bot) Open(ctx context.Context) error {
	if err := b.s.OpenGateway(ctx); err != nil {
		return err
	}

	_, err := b.s.Rest.SetGlobalCommands(b.s.ApplicationID, loader.Loader().Commands())
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) Close(ctx context.Context) {
	b.s.Close(ctx)
}
