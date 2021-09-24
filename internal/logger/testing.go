package logger

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

type TestLogRepository struct{}

func (t *TestLogRepository) Save(l *LogRequest) error {
	return nil
}

func (t *TestLogRepository) SaveRaw(uuid uuid.UUID, ip string, body []byte) error {
	return nil
}

func TestLogRequest(t *testing.T) *LogRequest {
	t.Helper()

	return &LogRequest{
		UserUUID:  uuid.New(),
		Timestamp: time.Now().Unix(),
		Events: []Events{
			{
				EventType: "log",
				EventText: "log message",
			},
		},
	}
}
