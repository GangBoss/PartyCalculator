package patterns

import (
	"github.com/GangBoss/PartyCalculator/src/types"
	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

type RequestTbAdapter struct {
	logger logrus.FieldLogger
}

func (m *RequestTbAdapter) Adapt(mess *tb.Message) *types.Request {
	m.logger.Infof("received message from tg chat: %v",mess.Chat.ID)
	return &types.Request{Text: mess.Text}
}