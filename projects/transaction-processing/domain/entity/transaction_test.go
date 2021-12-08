package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransaction_IsValid(t *testing.T) {
	// given
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 900

	// when
	err := transaction.IsValid()

	// then
	assert.Nil(t, err)
}

func TestTransaction_IsInvalidWithAmountGreaterThan1000(t *testing.T) {
	// given
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 1001

	// when
	err := transaction.IsValid()

	// then
	assert.Error(t, err)
	assert.Equal(t, "you dont have limit for this transaction", err.Error())
}

func TestTransaction_IsInvalidWithAmountLessThan1(t *testing.T) {
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
