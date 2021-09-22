package logger_test

import (
	"testing"

	"github.com/ZombieMInd/go-logger/internal/logger"
	"github.com/stretchr/testify/assert"
)

func TestLogRequest_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *logger.LogRequest
		isValid bool
	}{
		{
			name: "valid",
			u: func() *logger.LogRequest {
				return logger.TestLogRequest(t)
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
