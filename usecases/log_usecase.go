package usecases

import (
	"context"

	"github.com/saleamlakw/LoanTracker/domain/entities"
)

type logUserCase struct {
	logRepository entities.LogRepository
}

func NewLogUserCase(logRepository entities.LogRepository) *logUserCase {
	return &logUserCase{
		logRepository: logRepository,
	}
}

func (lu *logUserCase) GetAllLogs(ctx context.Context) ([]entities.Log, error) {
    return lu.logRepository.GetAllLogs(ctx)
}


