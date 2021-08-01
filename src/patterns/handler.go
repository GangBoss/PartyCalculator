package patterns

import (
	"github.com/GangBoss/PartyCalculator/src/types"
)

type Handler interface {
	HandleRequest(types.HandleResult) types.HandleResult
}
