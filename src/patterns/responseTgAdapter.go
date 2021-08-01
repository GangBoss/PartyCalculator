package patterns

import (
	"github.com/GangBoss/PartyCalculator/src/types"
	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

type ResponseTbAdapter struct {
	logger logrus.FieldLogger
}

func (m *ResponseTbAdapter) Adapt(result *types.HandleResult) (*tb.Sendable, string,tb.Recipient) {
	m.logger.Infof("processed message from tg chat")
	return nil, result.Response.Text,nil
}
