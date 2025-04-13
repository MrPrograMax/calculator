package service

import (
	myerror "calculator/internal/error"
	"calculator/internal/pkg/model"
	db "calculator/internal/pkg/storage"
	"context"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

func (s Service) Exec(ctx context.Context, expressions []*model.Expression) ([]*model.Result, error) {
	store := db.NewStorage()

	for _, expr := range expressions {
		switch expr.Type {
		case model.Calc:
			err := validateCalcRequest(expr)
			if err != nil {
				s.logger.Errorf("[pkg][service] validateCalcRequest: %v", err.Error())
				return nil, errors.Wrap(err, "validateCalcRequest")
			}

			err = s.Calc(ctx, store, expr)
			if err != nil {
				s.logger.Errorf("[pkg][service] Calc: %v", err.Error())
				return nil, errors.Wrap(err, "calculator.Calc")
			}
		case model.Print:
			err := s.Print(ctx, store, expr)
			if err != nil {
				s.logger.Errorf("[pkg][service] Print: %v", err.Error())
				return nil, errors.Wrap(err, "calculator.Print")
			}
		default:
			s.logger.Errorf("[pkg][service] %v", myerror.ErrUnknownOperation)
			return nil, errors.Wrap(myerror.ErrUnknownOperation, string(expr.Type))
		}
	}

	results, err := store.GetResults(ctx)
	if err != nil {
		s.logger.Errorf("[pkg][service]  GetResults: %v", err.Error())
		return nil, errors.Wrap(err, "store.GetResults")
	}

	return results, nil
}

func validateCalcRequest(req *model.Expression) error {
	var mErr error

	if req.Op == nil || *req.Op == "" {
		mErr = multierr.Append(mErr, errors.New("empty operation"))
	}

	if req.Left == nil || *req.Left == "" {
		mErr = multierr.Append(mErr, errors.New("empty left value"))
	}

	if req.Right == nil || *req.Right == "" {
		mErr = multierr.Append(mErr, errors.New("empty right value"))
	}

	return mErr
}
