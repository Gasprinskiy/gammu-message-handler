package repository

import (
	"tgsms/internal/entity/incoming_messages"
	"tgsms/internal/transaction"
)

type IncomingMessages interface {
	FindRecievedMessage(ts transaction.Session, sender string) (incoming_messages.Message, error)
	MarkMessageAsProcessed(ts transaction.Session, id int) error
}
