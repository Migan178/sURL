package url

import (
	"git.miganbox.com/migan/surl/configs"
	"git.miganbox.com/migan/surl/repository"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

func Create(d discord.SlashCommandInteractionData, e *handler.CommandEvent) error {
	url := d.String("url")

	createdData, err := repository.GetDatabase().CreateLink(e.Ctx, url)
	if err != nil {
		return err
	}

	_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateV2(
		discord.NewContainer(
			discord.NewTextDisplayf("### 단축 URL 생성 완료"),
			discord.NewTextDisplayf("- 단축 URL\n> `%s/%s`", configs.GetConfig().Bot.BackendURL, createdData.URN),
			discord.NewTextDisplayf("- 원본 URL\n> `%s`", createdData.RedirectURL),
			discord.NewTextDisplayf("- 생성된 날짜\n> %s", discord.FormattedTimestampMention(createdData.CreatedAt.Unix(), discord.TimestampStyleRelative)),
		),
	))
	return err
}
