package storage

import (
	myerror "calculator/internal/error"
	"calculator/internal/pkg/model"
	"context"
	"github.com/cockroachdb/errors"
	"strconv"
)

type Storage struct {
	variables map[string]int64
	result    []*model.Result
}

func NewStorage() *Storage {
	return &Storage{
		variables: make(map[string]int64),
	}
}

func (s *Storage) AddResult(ctx context.Context, variable string) error {
	if err := ctx.Err(); err != nil {
		return errors.Wrap(err, ctx.Err().Error())
	}

	s.result = append(s.result, &model.Result{
		Var: variable,
	})

	return nil
}

func (s *Storage) GetResults(ctx context.Context) ([]*model.Result, error) {
	if err := ctx.Err(); err != nil {
		return nil, errors.Wrap(err, ctx.Err().Error())
	}

	for _, result := range s.result {
		if _, ok := s.variables[result.Var]; ok {
			result.Value = s.variables[result.Var]
		} else {
			return nil, errors.Wrap(myerror.ErrUnknownVariable, result.Var)
		}
	}

	return s.result, nil
}

func (s *Storage) Action(ctx context.Context, operation model.Operation, variable, left, right string) error {
	if err := ctx.Err(); err != nil {
		return errors.Wrap(err, ctx.Err().Error())
	}

	leftValue, rightValue, err := s.getValues(ctx, left, right)
	if err != nil {
		return errors.Wrap(err, "s.getValues")
	}

	result, err := s.execOperation(operation, leftValue, rightValue)
	if err != nil {
		return errors.Wrap(err, "s.execOperation")
	}

	s.variables[variable] = result

	return nil
}

func (s *Storage) getValues(ctx context.Context, var1, var2 string) (int64, int64, error) {
	leftValue, err := s.tryToParseValue(ctx, var1)
	if err != nil {
		return 0, 0, errors.Wrap(err, "s.tryToParseValue")
	}

	rightValue, err := s.tryToParseValue(ctx, var2)
	if err != nil {
		return 0, 0, errors.Wrap(err, "s.tryToParseValue")
	}

	return leftValue, rightValue, nil
}

func (s *Storage) tryToParseValue(ctx context.Context, variable string) (int64, error) {
	if err := ctx.Err(); err != nil {
		return 0, errors.Wrap(err, ctx.Err().Error())
	}

	leftValue, err := strconv.ParseInt(variable, 10, 64)
	if err != nil {
		if v, ok := s.variables[variable]; ok {
			leftValue = v
		} else {
			return 0, errors.Wrap(myerror.ErrVariableNotFound, variable)
		}
	}

	return leftValue, nil
}

func (s *Storage) execOperation(operation model.Operation, v1, v2 int64) (int64, error) {
	switch operation {
	case model.Addition:
		return v1 + v2, nil
	case model.Subtraction:
		return v1 - v2, nil
	case model.Multiplication:
		return v1 * v2, nil
	case model.Division:
		if v2 == 0 {
			return v1, errors.Wrap(myerror.ErrDivideByZero, "execOperation")
		}

		return v1 / v2, nil
	}

	return 0, myerror.ErrUnknownOperation
}
