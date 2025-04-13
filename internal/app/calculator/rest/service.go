package rest

import (
	"calculator/internal/pkg/model"
	"context"
	"github.com/sirupsen/logrus"
)

type CalculatorREST interface {
	// Exec делает вычисления и возвращает результат.
	Exec(ctx context.Context, expressions []*model.Expression) ([]*model.Result, error)
}

type Implementation struct {
	logger         *logrus.Logger
	calculatorREST CalculatorREST
}

func NewRESTService(logger *logrus.Logger, calculator CalculatorREST) *Implementation {
	return &Implementation{
		logger:         logger,
		calculatorREST: calculator,
	}
}
