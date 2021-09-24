package teststore

import (
	"github.com/ZombieMInd/go-logger/internal/logger"
	"github.com/google/uuid"
)

type LogRepository struct {
	store *Store
	logs  map[uuid.UUID]*logger.LogRequest
}

func (r *LogRepository) Save(l *logger.LogRequest) error {
	r.logs[l.UUID] = l
	return nil
}

func (r *LogRepository) SaveRaw(uuid uuid.UUID, ip string, body []byte) error {
	return nil
}
