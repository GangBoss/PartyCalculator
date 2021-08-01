package handlers

import (
	"github.com/GangBoss/PartyCalculator/src/engine/handlers/executors"
	"github.com/sirupsen/logrus"
	"reflect"
)

// RouteHandler find executor by requestState
type RouteHandler struct {
	executors []executors.Executor
	logger    logrus.FieldLogger
}

func (m *RouteHandler) AddExecutor(executor executors.Executor) {
}

func NewRouter(logger *logrus.Logger) *RouteHandler {
	var router = &RouteHandler{}
	router.logger = logger.WithField("class", reflect.TypeOf(*router).String())
	router.logger.Infof("created")
	return router
}
