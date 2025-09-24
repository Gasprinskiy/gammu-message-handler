package usecase

import (
	"context"
	"tgsms/config"
	"tgsms/internal/entity/global"
	"tgsms/internal/entity/incoming_messages"
	"tgsms/internal/transaction"
	"tgsms/rimport"
	"tgsms/tools/logger"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/sirupsen/logrus"
)

type IncomingMessages struct {
	tgBot *bot.Bot
	conf  *config.Config
	log   *logger.Logger
	ri    *rimport.Repository
}

func NewIncomingMessages(
	tgBot *bot.Bot,
	conf *config.Config,
	log *logger.Logger,
	ri *rimport.Repository,
) *IncomingMessages {
	return &IncomingMessages{
		tgBot,
		conf,
		log,
		ri,
	}
}

func (u *IncomingMessages) OnMessageReceive(ctx context.Context, params incoming_messages.OnMessageReceiveParams) error {
	lf := logrus.Fields{
		"sender": params.SenderNumber,
	}

	ts := transaction.MustGetSession(ctx)
	message, err := u.ri.IncomingMessages.FindRecievedMessage(ts, params.SenderNumber)
	if err != nil {
		u.log.File.WithFields(lf).Errorln("ошибка при поиске сообщения:", err)
		return global.ErrInternalError
	}

	_, err = u.tgBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    u.conf.ChatID,
		ParseMode: models.ParseModeHTML,
		Text:      message.TgMessge(),
	})
	if err != nil {
		u.log.File.WithFields(lf).Errorln("не удалось отправить сообщение:", err)
		return global.ErrInternalError
	}

	if err = u.ri.IncomingMessages.MarkMessageAsProcessed(ts, message.ID); err != nil {
		u.log.File.WithFields(lf).Errorln("не удалось пометить сообщения как обработанное:", err)
		return global.ErrInternalError
	}

	u.log.File.WithFields(lf).Info("сообщение отправлено")
	return nil
}
