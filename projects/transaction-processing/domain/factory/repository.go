package factory

import "transaction-processing/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
