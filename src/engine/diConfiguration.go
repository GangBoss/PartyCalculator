package engine

import (
	"github.com/GangBoss/PartyCalculator/src/engine/handlers"
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
	container.Singleton(handlers.NewMainHandler)
	container.Singleton(handlers.NewAuthorizeHandler)
	configureMainHandler(container)
	return container
}

func configureMainHandler(container container.Container) {
	var mainHandler *handlers.MainHandler
	container.Make(&mainHandler)
	var authorizeHandler *handlers.AuthorizeHandler
	container.Make(&authorizeHandler)
	mainHandler.AddHandler(authorizeHandler)
}
