package repository

import (
	"aluno/adapter/repository/fixture"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTransactionRepositoryDb_Insert(t *testing.T) {
	// given
	migrationDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationDir)
	defer fixture.Down(db, migrationDir)

	repository := NewTransactionRepositoryDb(db)

	// when
	err := repository.Insert("1", "1", 2, "approved", "")

	// then
	assert.Nil(t, err)
}
