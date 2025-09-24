package mysql

import (
	"tgsms/internal/entity/incoming_messages"
	"tgsms/internal/repository"
	"tgsms/internal/transaction"
	"tgsms/tools/sql_gen"
)

type incomingMessages struct{}

func NewIncomingMessages() repository.IncomingMessages {
	return &incomingMessages{}
}

func (r *incomingMessages) FindRecievedMessage(ts transaction.Session, sender string) (incoming_messages.Message, error) {
	sqlQuery := `
		SELECT
			i.ID,
			i.ReceivingDateTime,
			i.TextDecoded,
			i.SenderNumber
		FROM inbox i
		WHERE i.SenderNumber = ?
			AND i.TextDecoded != ''
			AND i.Processed = 'false'
		ORDER BY i.ID DESC
		LIMIT 1
	`

	return sql_gen.Get[incoming_messages.Message](SqlxTx(ts), sqlQuery, sender)
}

func (r *incomingMessages) MarkMessageAsProcessed(ts transaction.Session, id int) error {
	sqlQuery := "UPDATE inbox SET Processed = 'true' WHERE ID = ?"

	_, err := SqlxTx(ts).Exec(sqlQuery, id)
	return err
}
