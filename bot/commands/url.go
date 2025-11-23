package commands

import (
	"github.com/Migan178/surl-bot/builders"
	"github.com/Migan178/surl-bot/commands/subcommands/url"
	"github.com/bwmarrin/discordgo"
)

var (
	urlCommandGet    = "확인"
	urlCommandCreate = "생성"
)

var URLCommand = &Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "url",
		Description: "asdf",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        urlCommandGet,
				Description: "해당 단축 URL의 정보를 확인해요.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "url",
						Description: "정보를 확인할 단축 URL를 입력해 주세요.",
						Required:    true,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        urlCommandCreate,
				Description: "단축 URL를 생성해요.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "url",
						Description: "단축할 URL를 입력해 주세요.",
						Required:    true,
					},
				},
			},
		},
	},
	Run: func(inter *builders.InteractionCreate) error {
		switch opt := inter.ApplicationCommandData().Options[0]; opt.Name {
		case urlCommandCreate:
			if err := inter.DeferReply(&discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
			}); err != nil {
				return err
			}

			return url.Create(inter, opt.Options[0].StringValue())
		case urlCommandGet:
			if err := inter.DeferReply(&discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
			}); err != nil {
				return err
			}

			return url.Get(inter, opt.Options[0].StringValue())
		default:
			return nil
		}
	},
}

func init() {
	GetDiscommand().LoadCommand(URLCommand)
}
