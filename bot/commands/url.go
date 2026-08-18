package commands

import (
	"git.miganbox.com/migan/surl/bot/commands/subcommands/url"
	"git.miganbox.com/migan/surl/bot/loader"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/disgo/handler/middleware"
)

func init() {
	const (
		urlCommandGet    = "확인"
		urlCommandCreate = "생성"
	)

	loader.Loader().RegisterCommand(discord.SlashCommandCreate{
		Name:        "url",
		Description: "asdf",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionSubCommand{
				Name:        urlCommandCreate,
				Description: "단축 URL를 생성해요.",
				Options: []discord.ApplicationCommandOption{
					discord.ApplicationCommandOptionString{
						Name:        "url",
						Description: "단축할 URL를 입력해 주세요.",
						Required:    true,
					},
				},
			},
			discord.ApplicationCommandOptionSubCommand{
				Name:        urlCommandGet,
				Description: "해당 단축 URL의 정보를 확인해요.",
				Options: []discord.ApplicationCommandOption{
					discord.ApplicationCommandOptionString{
						Name:        "url",
						Description: "정보를 확인할 단축 URL를 입력해 주세요.",
						Required:    true,
					},
				},
			},
		},
	})

	loader.Loader().RegisterHandler(func(r handler.Router) {
		r.Use(middleware.Defer(discord.InteractionTypeApplicationCommand, false, true))

		r.Route("/url", func(r handler.Router) {
			r.SlashCommand("/"+urlCommandCreate, url.Create)
			r.SlashCommand("/"+urlCommandGet, url.Get)
		})
	})
}
