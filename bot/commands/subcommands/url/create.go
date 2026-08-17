package url

import (
	"context"
	"fmt"
	"time"

	"git.miganbox.com/migan/surl/bot/builders"
	"git.miganbox.com/migan/surl/configs"
	"git.miganbox.com/migan/surl/repository"
)

func Create(inter *builders.InteractionCreate, url string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	createdData, err := repository.GetDatabase().CreateLink(ctx, url)
	if err != nil {
		return err
	}

	return builders.NewMessageSender(inter).
		AddComponents(
			builders.ContainerBuilder().
				AddText("### 단축 URL 생성 완료").
				AddText(fmt.Sprintf("- 단축 URL\n> `%s/%s`", configs.GetConfig().Bot.BackendURL, createdData.URN)).
				AddText(fmt.Sprintf("- 원본 URL\n> `%s`", createdData.RedirectURL)).
				AddText(fmt.Sprintf("- 생성된 날짜\n> %s", builders.Time(&createdData.CreatedAt, builders.RelativeTime))),
		).
		SetComponentsV2(true).
		SetEphemeral(true).
		Send()
}
