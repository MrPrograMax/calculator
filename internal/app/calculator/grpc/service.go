package grpc

import (
	"calculator/internal/pkg/model"
	desc "calculator/pkg/api"
	"context"
	"github.com/sirupsen/logrus"
)

type Calculator interface {
	// Exec делает вычисления и возвращает результат.
	Exec(ctx context.Context, expressions []*model.Expression) ([]*model.Result, error)
}

type Implementation struct {
	logger *logrus.Logger

	calculator Calculator
	desc.UnimplementedURLServiceServer
}

func NewGRPSService(logger *logrus.Logger, calculator Calculator) *Implementation {
	return &Implementation{
		logger:     logger,
		calculator: calculator,
	}
}
