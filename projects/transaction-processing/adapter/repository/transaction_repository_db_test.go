package repository

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"transaction-processing/adapter/repository/fixture"
)

func TestTransactionRepositoryDb_Insert(t *testing.T) {
	// given
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepositoryDb(db)

	// when
	err := repository.Insert("1", "1", 2, "approved", "")

	// then
	assert.Nil(t, err)
}
