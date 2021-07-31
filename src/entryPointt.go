package main

import (
	"github.com/GangBoss/PartyCalculator/src/engine"
	"log"
)

func main() {

var container= engine.ConfigureDi()
	var logLink *log.Logger
	container.Make(&logLink)
	var logInst *log.Logger
	container.Make(&logInst)
	logLink.Print("log")
	logInst.Print("log2")


}
