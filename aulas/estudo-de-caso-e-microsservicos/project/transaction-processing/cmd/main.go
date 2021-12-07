package main

import (
	"database/sql"
	"encoding/json"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"transaction-processing/adapter/broker/kafka"
	"transaction-processing/adapter/factory"
	"transaction-processing/adapter/presenter/transaction"
	"transaction-processing/usecase/process_transaction"
)

func main() {
	// db
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	// repository
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()

	// configMapProducer
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}

	// producer
	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	var msgChan = make(chan *ckafka.Message)

	// configMapConsumer
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp",
		"group.id":          "goapp",
	}
	topics := []string{"transactions"}

	// consumer
	consumer := kafka.NewConsumer(configMapConsumer, topics)

	// topic
	go consumer.Consume(msgChan) // criar uma Thread separada para consumir o tópico infinitamente

	// usecase
	usecase := process_transaction.NewProcessTransaction(repository, producer, "transactions_result")

	for msg := range msgChan {
		var input process_transaction.TransactionDtoInput
		json.Unmarshal(msg.Value, &input)
		usecase.Execute(input)
	}
}
