package executors

import (
	"github.com/GangBoss/PartyCalculator/src/types"
	"github.com/sirupsen/logrus"
	"reflect"
)

type RoomCommandExecutor struct {
	logger logrus.FieldLogger
}

func (m *RoomCommandExecutor) HandleRequest(result types.HandleResult) types.HandleResult {
	return result
}

func NewCommandExecutor(logger *logrus.Logger) *RoomCommandExecutor {
	var authorizeHandler = new(RoomCommandExecutor)
	authorizeHandler.logger = logger.WithField("class", reflect.TypeOf(*authorizeHandler).String())
	authorizeHandler.logger.Infof("created")
	return authorizeHandler
}
