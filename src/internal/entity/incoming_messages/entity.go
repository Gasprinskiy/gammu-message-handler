package incoming_messages

import (
	"fmt"
	"tgsms/tools/chronos"
	"time"
)

type Message struct {
	ID          int       `db:"ID"`
	ReceiveTime time.Time `db:"ReceivingDateTime"`
	Sender      string    `db:"SenderNumber"`
	Text        string    `db:"TextDecoded"`
}

func (m Message) TgMessge() string {
	return fmt.Sprintf(
		incomingMessageTemplate,
		m.Sender,
		m.ReceiveTime.Format(chronos.DateTimeMask),
		m.Text,
	)
}
