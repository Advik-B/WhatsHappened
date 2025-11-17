package models

import "time"

type ParsedMessage struct {
	Time    time.Time
	Sender  string
	Content string
}
