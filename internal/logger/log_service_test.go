package logger_test

import (
	"testing"

	"github.com/ZombieMInd/go-logger/internal/logger"
	"github.com/stretchr/testify/assert"
)

func TestLogService_Save(t *testing.T) {
	service := logger.NewService(&logger.TestLogRepository{})

	assert.NoError(t, service.Save(logger.TestLogRequest(t)))
}
