package redisstore

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ZombieMInd/go-logger/internal/logger"
	"github.com/google/uuid"
)

type LogRepository struct {
	store *Store
}

func (r *LogRepository) Save(l *logger.LogRequest) error {
	key := fmt.Sprintf("%s:%s", l.UserUUID, l.UUID.String())

	jsonLog, err := json.Marshal(l)
	if err != nil {
		return err
	}

	err = r.store.client.Set(key, jsonLog, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *LogRepository) SaveRaw(uuid uuid.UUID, ip string, body []byte) error {
	key := fmt.Sprintf("%s:%s:%s", uuid, ip, time.Now())
	return r.store.client.Set(key, body, 0).Err()
}
