package handlers

import (
	"github.com/GangBoss/PartyCalculator/src/types"
	"github.com/sirupsen/logrus"
	"reflect"
)

type StateHandler struct {
	logger logrus.FieldLogger
}
//HandleRequest sets command to handle result
func (m *StateHandler) HandleRequest(result types.HandleResult) types.HandleResult {
	return result
}

func NewStateHandler(logger *logrus.Logger) *StateHandler {
	var stateHandler = new(StateHandler)
	stateHandler.logger = logger.WithField("class", reflect.TypeOf(*stateHandler).String())
	stateHandler.logger.Infof("created")
	return stateHandler
}
