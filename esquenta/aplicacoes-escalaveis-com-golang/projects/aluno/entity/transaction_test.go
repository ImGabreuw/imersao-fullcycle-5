package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransactionWithAmountGreaterThan1000(t *testing.T) {
	// given
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 2000

	// when
	err := transaction.IsValid()

	// then
	assert.Error(t, err)
	assert.Equal(t, "you dont have limit for this transaction", err.Error())
}

func TestTransactionWithAmountLessThan1(t *testing.T) {
	// given
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0

	// when
	err := transaction.IsValid()

	// then
	assert.Error(t, err)
	assert.Equal(t, "the amount must be greater than 1", err.Error())
}
