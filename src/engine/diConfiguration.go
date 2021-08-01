package engine

import (
	"github.com/golobby/container/pkg/container"
	"github.com/sirupsen/logrus"
	"os"
)

func ConfigureDi() container.Container {
	container := container.NewContainer()


	container.Singleton(func ()*logrus.Logger{
		log:=logrus.New()
		log.SetFormatter(&logrus.JSONFormatter{})
		log.Out=os.Stdout
		return log
	})
	container.Singleton(NewMainHandler)
	container.Singleton(NewAuthorizeHandler)
	configureMainHandler(container)
	return container
}

func configureMainHandler(container container.Container) {
	var mainHandler *MainHandler
	container.Make(&mainHandler)
	var authorizeHandler *AuthorizeHandler
	container.Make(&authorizeHandler)
	mainHandler.AddHandler(authorizeHandler)
}
