package patterns

import "github.com/GangBoss/PartyCalculator/src/types"

type Statable interface {
	GetRequestState(request types.Request) string
}