package engine

import (
	"github.com/GangBoss/PartyCalculator/src/patterns"
	"github.com/GangBoss/PartyCalculator/src/types"
	"github.com/sirupsen/logrus"
	"reflect"
)

type MainHandler struct {
	handlers []patterns.Handler
	logger   logrus.FieldLogger
}

func (m *MainHandler) HandleRequest(result types.HandleResult, observer chan<- *types.HandleResult) {
	for _, handler := range m.handlers {
		result = (handler).HandleRequest(result)
		if !result.Success() {
			m.logger.Errorf("Handle result failed on %v with %v ", handler, result)
		}

		if !result.Continuable {
			m.logger.Infof("Handle result success result :%v ", result)
			break
		}
	}
	observer <- &result
}

func (m *MainHandler) AddHandler(handler patterns.Handler) {
	m.logger.Infof("handler %s added to pipeline",reflect.TypeOf(handler).String())

	m.handlers = append(m.handlers, handler)
}

func (m *MainHandler) WaitForRequest(initialRequest <-chan *types.HandleResult, observer chan<- *types.HandleResult) {
	defer func() {}()
	{
		err := recover()
		if err != nil {
			m.logger.Errorf("caught an error %s", err)
		}
	}
	for val := range initialRequest {
		go m.HandleRequest(*val, observer)
	}
}

func NewMainHandler(logger *logrus.Logger) *MainHandler {
	var mainHandler = &MainHandler{ handlers: nil }
	mainHandler.logger = logger.WithField("class", reflect.TypeOf(*mainHandler).String())
	mainHandler.logger.Info("created")
	return mainHandler
}
