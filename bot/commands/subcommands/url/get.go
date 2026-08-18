package url

import (
	"git.miganbox.com/migan/surl/repository"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

func Get(d discord.SlashCommandInteractionData, e *handler.CommandEvent) error {
	url := d.String("url")

	data, err := repository.GetDatabase().Find(e.Ctx, url)
	if err != nil {
		return err
	}

	_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateV2(
		discord.NewContainer(
			discord.NewTextDisplayf("### 해당 단축 URL의 정보"),
			discord.NewTextDisplayf("- 단축 URL\n> `%s`", url),
			discord.NewTextDisplayf("- 원본 URL\n> `%s`", data.RedirectURL),
			discord.NewTextDisplayf("- 생성된 날짜\n> %s", discord.FormattedTimestampMention(data.CreatedAt.Unix(), discord.TimestampStyleRelative)),
		),
	))
	return err
}
