package logger

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type LogRequest struct {
	UUID      uuid.UUID
	IP        string
	UserUUID  uuid.UUID `json:"user_uuid"`
	Timestamp int64     `json:"timestamp"`
	Events    []Events
}

type Events struct {
	EventType string `json:"event_type"`
	EventText string `json:"event_txt"`
}

func (l *LogRequest) Validate() error {
	return validation.ValidateStruct(
		l,
		validation.Field(&l.UserUUID, is.UUID),
	)
}
