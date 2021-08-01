package main

import (
	"github.com/GangBoss/PartyCalculator/src/engine"
	"github.com/sirupsen/logrus"
)

func main() {

var container= engine.ConfigureDi()
	var logLink *logrus.Logger
	container.Make(&logLink)
	logLink.Printf("started")


}
