package url

import (
	"fmt"

	"github.com/Migan178/surl-bot/builders"
	"github.com/Migan178/surl-bot/client"
	"github.com/Migan178/surl-bot/configs"
)

func Create(inter *builders.InteractionCreate, url string) error {
	createdData, err := client.GetClient().Create(url)
	if err != nil {
		return err
	}

	return builders.NewMessageSender(inter).
		AddComponents(
			builders.ContainerBuilder().
				AddText("### 단축 URL 생성 완료").
				AddText(fmt.Sprintf("- 단축 URL\n> `%s/%s`", configs.GetConfig().SURL.URL, createdData.URN)).
				AddText(fmt.Sprintf("- 원본 URL\n> `%s`", createdData.RedirectURL)).
				AddText(fmt.Sprintf("- 생성된 날짜\n> %s", builders.Time(&createdData.CreatedAt, builders.RelativeTime))),
		).
		SetComponentsV2(true).
		SetEphemeral(true).
		Send()
}
