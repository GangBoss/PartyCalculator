package engine

import (
	"container/list"
	"github.com/GangBoss/PartyCalculator/src/types"
)
type AuthorizeHandler struct {
	handlers *list.List
}

func (m *AuthorizeHandler) HandleRequest(result types.HandleResult, observer chan *types.HandleResult) {
	observer<-types.NewHandleResult()
}
func NewAuthorizeHandler() *AuthorizeHandler {
	var authorizeHandler = new(AuthorizeHandler)
	return authorizeHandler
}