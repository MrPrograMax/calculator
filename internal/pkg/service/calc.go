package service

import (
	"calculator/internal/pkg/model"
	"calculator/internal/pkg/storage"
	"context"
)

func (s Service) Calc(ctx context.Context, storage *storage.Storage, expression *model.Expression) error {
	return storage.Action(ctx, *expression.Op, expression.Var, *expression.Left, *expression.Right)
}
