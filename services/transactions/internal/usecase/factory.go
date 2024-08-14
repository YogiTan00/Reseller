package usecase

import (
	"github.com/YogiTan00/Reseller/pkg/logger"
	"github.com/YogiTan00/Reseller/services/transactions/domain/repository"
	"github.com/YogiTan00/Reseller/services/transactions/domain/usecase"
)

type TransactionUsecase struct {
	l               logger.Logger
	repoTransaction repository.TransactionRepositoryInterface
}

type TransactionUsecaseFactory struct {
	L               logger.Logger
	RepoTransaction repository.TransactionRepositoryInterface
}

func (prod *TransactionUsecaseFactory) Create() usecase.TransactionUsecaseInterface {
	return &TransactionUsecase{
		l:               prod.L,
		repoTransaction: prod.RepoTransaction,
	}
}
