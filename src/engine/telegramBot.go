package engine

import (
	"github.com/GangBoss/PartyCalculator/src/engine/handlers"
	"github.com/GangBoss/PartyCalculator/src/patterns"
	"github.com/GangBoss/PartyCalculator/src/types"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

func startBot(handler handlers.MainHandler, adapterFrom patterns.RequestTbAdapter, adapterTo patterns.ResponseTbAdapter) {
	b, err := tb.NewBot(tb.Settings{
		URL:    "https://api.telegram.org",
		Token:  "TOKEN_HERE",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}
	initialChan := make(chan *types.Request)
	observerChan := make(chan *types.HandleResult)
	handler.WaitForRequest(initialChan, observerChan)

	b.Handle("", func(m *tb.Message) {
		req := adapterFrom.Adapt(m)
		initialChan <- req
	})

	b.Start()

	for val := range observerChan {
		sendable, mess, recipient := adapterTo.Adapt(val)
		if sendable != nil {
			b.Send(recipient, sendable)
		}
		b.Send(recipient, mess)
	}
}
