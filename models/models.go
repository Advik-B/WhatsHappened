package models

import (
	"fmt"
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

func MakeUserSender(displayName string) MessageSender {
	return MessageSender{
		Type:    MessageSenderTypeUser,
		display: &displayName,
	}
}

func (m *MessageSender) BuildMessage(content []byte, msgType MessageType, timestamp time.Time, originalFileName *string) (Message, error) {
	switch msgType {
	case MessageTypeText:
		return MakeTextMessage(string(content), timestamp), nil
	case MessageTypeImage:
		return MakeImageMessage(*originalFileName, content, timestamp), nil
	case MessageTypeVideo:
		return MakeVideoMessage(*originalFileName, content, timestamp), nil
	case MessageTypeAudio:
		return MakeAudioMessage(*originalFileName, content, timestamp), nil
	default:
		return Message{}, fmt.Errorf("a")
	}
}

type Message struct {
	TimeStamp        time.Time
	Type             MessageType
	Content          []byte  // Content can be text or binary data depending on Type
	OriginalFileName *string // Optional: original file name for media messages
	Sender           MessageSender
}

func MakeTextMessage(text string, timestamp time.Time) Message {
	return Message{
		TimeStamp:        timestamp,
		Type:             MessageTypeText,
		Content:          []byte(text),
		OriginalFileName: nil,
	}
}

func MakeImageMessage(filename string, imageData []byte, timestamp time.Time) Message {
	return Message{
		TimeStamp:        timestamp,
		Type:             MessageTypeImage,
		Content:          imageData,
		OriginalFileName: &filename,
	}
}
func MakeVideoMessage(filename string, videoData []byte, timestamp time.Time) Message {
	return Message{
		TimeStamp:        timestamp,
		Type:             MessageTypeVideo,
		Content:          videoData,
		OriginalFileName: &filename,
	}
}

func MakeAudioMessage(filename string, audioData []byte, timestamp time.Time) Message {
	return Message{
		TimeStamp:        timestamp,
		Type:             MessageTypeAudio,
		Content:          audioData,
		OriginalFileName: &filename,
	}
}
