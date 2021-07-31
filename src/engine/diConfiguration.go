package engine

import (
	"github.com/golobby/container/pkg/container"
	"log"
	"os"
)

func ConfigureDi() container.Container {
	container := container.NewContainer()
	container.Transient(func() *log.Logger {
		return log.New(os.Stdout, " ", 0)
	})
	container.Singleton(NewMainHandler)
	var handler *MainHandler
	container.Make(&handler)
	return container
}
