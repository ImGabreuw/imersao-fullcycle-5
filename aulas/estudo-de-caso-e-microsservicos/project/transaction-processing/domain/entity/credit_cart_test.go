package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreditCard_NumberIsInvalid(t *testing.T) {
	// given
	number := "40000000000000000"
	name := "Jose da Silva"
	expirationMonth := 12
	expirationYear := 2024
	cvv := 123

	// when
	_, err := NewCreditCard(number, name, expirationMonth, expirationYear, cvv)

	// then
	assert.Equal(t, "invalid credit card number", err.Error())
}

func TestCreditCard_NumberIsValid(t *testing.T) {
	// given
	number := "4193523830170205"
	name := "Jose da Silva"
	expirationMonth := 12
	expirationYear := 2024
	cvv := 123

	// when
	_, err := NewCreditCard(number, name, expirationMonth, expirationYear, cvv)

	// then
	assert.Nil(t, err)
}

func TestCreditCard_ExpirationMonthIsInvalidBecauseExpirationMonthIs13(t *testing.T) {
	// given
	number := "4193523830170205"
	name := "Jose da Silva"
	expirationMonth := 13
	expirationYear := 2024
	cvv := 123

	// when
	_, err := NewCreditCard(number, name, expirationMonth, expirationYear, cvv)

	// then
	assert.Equal(t, "invalid expiration month", err.Error())
}

func TestCreditCard_ExpirationMonthIsInvalidBecauseExpirationMonthIs0(t *testing.T) {
	// given
	number := "4193523830170205"
	name := "Jose da Silva"
	expirationMonth := 0
	expirationYear := 2024
	cvv := 123

	// when
	_, err := NewCreditCard(number, name, expirationMonth, expirationYear, cvv)

	// then
	assert.Equal(t, "invalid expiration month", err.Error())
}

func TestCreditCard_ExpirationMonthIsValid(t *testing.T) {
	// given
	number := "4193523830170205"
	name := "Jose da Silva"
	expirationMonth := 11
	expirationYear := 2024
	cvv := 123

	// when
	_, err := NewCreditCard(number, name, expirationMonth, expirationYear, cvv)

	// then
	assert.Nil(t, err)
}

func TestCreditCard_ExpirationYearIsInvalid(t *testing.T) {
	// given
	number := "4193523830170205"
	name := "Jose da Silva"
	expirationMonth := 11
	lastYear := time.Now().AddDate(-1, 0, 0)
	expirationYear := lastYear.Year()
	cvv := 123

	// when
	_, err := NewCreditCard(number, name, expirationMonth, expirationYear, cvv)

	// then
	assert.Equal(t, "invalid expiration year", err.Error())
}
