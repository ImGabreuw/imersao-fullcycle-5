package main

import (
	"database/sql"
	"encoding/json"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"transaction-processing/adapter/broker/kafka"
	"transaction-processing/adapter/factory"
	"transaction-processing/adapter/presenter/transaction"
	"transaction-processing/usecase/process_transaction"
)

func main() {
	// db
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USERNAME")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST")+":3306)/"+os.Getenv("MYSQL_DATABASE"))

	if err != nil {
		log.Fatal(err)
	}

	// repository
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()

	// configMapProducer
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("BOOTSTRAP_SERVERS"),
		"security.protocol": os.Getenv("SECURITY_PROTOCOL"),
		"sasl.mechanisms":   os.Getenv("SASL_MECHANISMS"),
		"sasl.username":     os.Getenv("SASL_USERNAME"),
		"sasl.password":     os.Getenv("SASL_PASSWORD"),
	}

	// producer
	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	var msgChan = make(chan *ckafka.Message)

	// configMapConsumer
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("BOOTSTRAP_SERVERS"),
		"security.protocol": os.Getenv("SECURITY_PROTOCOL"),
		"sasl.mechanisms":   os.Getenv("SASL_MECHANISMS"),
		"sasl.username":     os.Getenv("SASL_USERNAME"),
		"sasl.password":     os.Getenv("SASL_PASSWORD"),
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
