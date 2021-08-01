package handlers

import (
	"github.com/GangBoss/PartyCalculator/src/patterns"
	"github.com/GangBoss/PartyCalculator/src/types"
	"github.com/sirupsen/logrus"
	"reflect"
)

// AuthorizeHandler first handler to added, find user from user repository
type AuthorizeHandler struct {
	logger logrus.FieldLogger
	userRepository patterns.Repository
}

func (m *AuthorizeHandler) HandleRequest(result types.HandleResult) types.HandleResult {
	return result
}

func NewAuthorizeHandler(logger *logrus.Logger, userRepository patterns.Repository) *AuthorizeHandler {
	var authorizeHandler = &AuthorizeHandler{userRepository: userRepository}
	authorizeHandler.logger = logger.WithField("class", reflect.TypeOf(*authorizeHandler).String())
	authorizeHandler.logger.Infof("created")
	return authorizeHandler
}
