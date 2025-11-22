package url

import (
	"fmt"

	"github.com/Migan178/surl-bot/builders"
	"github.com/Migan178/surl-bot/client"
)

func Get(inter *builders.InteractionCreate, url string) error {
	data, err := client.GetClient().GetInformation(url)
	if err != nil {
		return err
	}

	return builders.NewMessageSender(inter).
		AddComponents(
			builders.ContainerBuilder().
				AddText("### 해당 단축 URL의 정보").
				AddText(fmt.Sprintf("- 단축 URL\n> `%s`", url)).
				AddText(fmt.Sprintf("- 원본 URL\n> `%s`", data.RedirectURL)).
				AddText(fmt.Sprintf("- 생성된 날짜\n> %s", builders.Time(&data.CreatedAt, builders.RelativeTime))),
		).
		SetComponentsV2(true).
		SetEphemeral(true).
		Send()
}
