package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
	"testing"
	"transaction-processing/adapter/presenter/transaction"
	"transaction-processing/domain/entity"
	"transaction-processing/usecase/process_transaction"
)

func TestProducer_Publish(t *testing.T) {
	// given
	expectedOutput := process_transaction.TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you dont have limit for this transaction",
	}

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}

	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())

	// when
	err := producer.Publish(expectedOutput, []byte("1"), "test")

	// then
	assert.Nil(t, err)
}
