package engine

import (
	"github.com/GangBoss/PartyCalculator/src/types"
	"github.com/sirupsen/logrus"
	"reflect"
)

type AuthorizeHandler struct {
	logger logrus.FieldLogger
}

func (m *AuthorizeHandler) HandleRequest(result types.HandleResult) types.HandleResult {
	return result
}

func NewAuthorizeHandler(logger *logrus.Logger) *AuthorizeHandler {
	var authorizeHandler = new(AuthorizeHandler)
	authorizeHandler.logger = logger.WithField("class", reflect.TypeOf(*authorizeHandler).String())
	authorizeHandler.logger.Infof("created")
	return authorizeHandler
}
