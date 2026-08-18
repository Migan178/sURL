package loader

import (
	"sync"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

type LoaderStruct struct {
	commands []discord.ApplicationCommandCreate
	r        handler.Router
}

var instance *LoaderStruct
var once sync.Once

func Loader() *LoaderStruct {
	once.Do(func() {
		instance = &LoaderStruct{r: handler.New()}
	})

	return instance
}

func (l *LoaderStruct) RegisterCommand(command discord.ApplicationCommandCreate) {
	l.commands = append(l.commands, command)
}

func (l *LoaderStruct) Commands() []discord.ApplicationCommandCreate {
	commandsCopy := make([]discord.ApplicationCommandCreate, len(l.commands))
	copy(commandsCopy, l.commands)
	return commandsCopy
}

func (l *LoaderStruct) Router() handler.Router {
	return l.r
}

func (l *LoaderStruct) RegisterHandler(f func(r handler.Router)) {
	l.r.Group(func(r handler.Router) {
		f(r)
	})
}
