package engine

import (
	"github.com/GangBoss/PartyCalculator/src/patterns"
	"github.com/GangBoss/PartyCalculator/src/types"
	"log"
	"reflect"
)

type MainHandler struct {
	handlers []patterns.Handler
	logger   *log.Logger
}

func (m *MainHandler) HandleRequest(result *types.HandleResult, observer chan<- *types.HandleResult) {
	for _, handler := range m.handlers {
		result = handler.HandleRequest(result)
		if !result.Success() {
			m.logger.Fatalf("Handle result failed on %v with %v ", handler, result)
			break
		}

		if !result.Continuable {
			m.logger.Printf("Handle result success result :%v ", result)

			break
		}
	}
	observer <- result
}

func (m *MainHandler) AddHandler(handler patterns.Handler) {
	m.handlers = append(m.handlers, handler)
}

func (m *MainHandler) WaitForRequest(initialRequest <-chan *types.HandleResult, observer chan<- *types.HandleResult) {
	defer func() {}()
	{
		error := recover()
		if error != nil {
			m.logger.Fatalf("caught an error %s ", error)
		}
	}
	for val := range initialRequest {
		go m.HandleRequest(val, observer)
	}
}

func NewMainHandler(logger *log.Logger) *MainHandler {
	var mainHandler = new(MainHandler)
	logger.SetPrefix(reflect.TypeOf(mainHandler).String() + " : ")
	mainHandler.logger = logger
	return mainHandler
}
