package executors

import (
	"github.com/GangBoss/PartyCalculator/src/types"
)

// Executor statablePattern to execute command
type Executor interface {
	CanExecute(command types.Command) bool
	Handle(command types.HandleResult) types.HandleResult
}


