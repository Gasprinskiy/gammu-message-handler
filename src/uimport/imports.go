package uimport

import (
	"tgsms/config"
	"tgsms/internal/usecase"
	"tgsms/rimport"
	"tgsms/tools/logger"

	"github.com/go-telegram/bot"
)

type Usecase struct {
	IncomingMessages *usecase.IncomingMessages
}

func NewUsecaseImport(
	ri *rimport.Repository,
	log *logger.Logger,
	conf *config.Config,
	b *bot.Bot,
) *Usecase {
	return &Usecase{
		IncomingMessages: usecase.NewIncomingMessages(b, conf, log, ri),
	}
}
