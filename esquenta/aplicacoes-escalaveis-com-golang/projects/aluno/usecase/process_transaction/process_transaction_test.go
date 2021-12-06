package process_transaction

import (
	mock_entity "aluno/entity/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessTransactionWhenItIsValid(t *testing.T) {
	// given
	input := TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    200,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       "approved",
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_entity.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, "approved", "").Return(nil)

	// when
	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)

	// then
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
