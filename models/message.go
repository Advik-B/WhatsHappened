package models

import (
	"time"
)

type MessageType int

const (
	MessageTypeText MessageType = iota
	MessageTypeImage
	MessageTypeVideo
	MessageTypeAudio
)

type MessageSenderType int

const (
	MessageSenderTypeUser MessageSenderType = iota
	MessageSenderTypeSystem
)

type MessageSender struct {
	Type    MessageSenderType
	display *string // Optional: display name for the sender
}

type Message struct {
	TimeStamp        time.Time
	Type             MessageType
	Content          []byte  // Content can be text or binary data depending on Type
	OriginalFileName *string // Optional: original file name for media messages
	Sender           MessageSender
}
