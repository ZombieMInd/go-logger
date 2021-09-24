package logger

import (
	"github.com/google/uuid"
)

type LogService struct {
	repository LogRepository
}

type LogRepository interface {
	Save(*LogRequest) error
	SaveRaw(uuid.UUID, string, []byte) error
}

func NewService(r LogRepository) *LogService {
	return &LogService{repository: r}
}

func (s *LogService) Save(l *LogRequest) error {
	l.UUID = uuid.New()
	err := s.repository.Save(l)
	return err
}

func (s *LogService) SaveRaw(l []byte, ip string) error {
	uuid := uuid.New()
	return s.repository.SaveRaw(uuid, ip, l)
}
