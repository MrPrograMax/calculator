package grpc

import (
	myerror "calculator/internal/error"
	"calculator/internal/pkg/converter"
	desc "calculator/pkg/api"
	"context"
)

func (i *Implementation) Calculate(ctx context.Context, req *desc.CalculateRequest) (*desc.CalculateResponse, error) {
	expressions := converter.ExpressionToModel(req)

	results, err := i.calculator.Exec(ctx, expressions)
	if err != nil {
		i.logger.Errorf("[app][calculator][grpc] Exec: %v", err.Error())
		return nil, myerror.ErrSomethingWentWrong
	}

	response := converter.ResultToProto(results)

	return response, nil
}
