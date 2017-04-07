package qtypes

import (
	msgTypes "github.com/qnib/qframe/types"
)


// Message type holds all the information for a single log event
type Message struct {
	msgTypes.QMsg
	Name       string            `json:"name"`
}

// NewMessage returns a new message
func NewMessage(typ, source, name string) Message {
	return Message {
		QMsg: msgTypes.NewQMsg(typ, source),
		Name:       name,
	}
}
