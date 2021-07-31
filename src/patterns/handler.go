package patterns

import "github.com/GangBoss/PartyCalculator/src/types"

type Handler interface {
	HandleRequest(result *types.HandleResult) *types.HandleResult
}
