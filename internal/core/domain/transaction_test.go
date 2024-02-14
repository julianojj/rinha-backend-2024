package domain

import (
	"testing"

	"github.com/julianojj/rinha-backend-2024/internal/core/exception"
	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	transaction := NewTransaction(1, 1000, "c", "Monitor")
	assert.NoError(t, transaction.Validate())
	assert.Equal(t, int64(1000), transaction.Amount)
	assert.Equal(t, "c", transaction.Type)
	assert.Equal(t, "Monitor", transaction.Description)
}

func TestReturnExceptionIfInvalidAmount(t *testing.T) {
	transaction := NewTransaction(1, -1000, "c", "Monitor")
	assert.EqualError(t, transaction.Validate(), exception.INVALID_AMOUNT)
}

func TestReturnExceptionIfInvalidDescription(t *testing.T) {
	transaction := NewTransaction(1, 1000, "c", "Monitor Dell")
	assert.EqualError(t, transaction.Validate(), exception.INVALID_DESCRIPTION)
}
