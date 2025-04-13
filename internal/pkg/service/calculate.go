package service

import (
	"calculator/internal/pkg/model"
	"context"
	"github.com/cockroachdb/errors"
)

func (s Service) Calculate(ctx context.Context, expressions []*model.Expression) ([]*model.Result, error) {
	results, err := s.Exec(ctx, expressions)
	if err != nil {
		s.logger.Errorf("[pkg][service][calculate] Error: %v", err.Error())
		return nil, errors.Wrap(err, "Exec")
	}

	return results, nil
}
