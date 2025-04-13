package service

import (
	"calculator/internal/pkg/model"
	"calculator/internal/pkg/storage"
	"context"
)

func (s Service) Print(ctx context.Context, storage *storage.Storage, expression *model.Expression) error {
	return storage.AddResult(ctx, expression.Var)
}
