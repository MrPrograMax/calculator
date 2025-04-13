package converter

import (
	"calculator/internal/pkg/model"
	desc "calculator/pkg/api"
	"github.com/samber/lo"
	"strconv"
)

func ExpressionToModel(request *desc.CalculateRequest) []*model.Expression {
	var expressions []*model.Expression

	for _, expr := range request.Data {
		expressions = append(expressions, &model.Expression{
			Type:  convertTypeToModel(expr.Type),
			Op:    convertOperationToModel(expr.Op),
			Var:   expr.Var,
			Left:  expr.Left,
			Right: expr.Right,
		})
	}

	return expressions
}

func ResultToProto(result []*model.Result) *desc.CalculateResponse {
	if result == nil || len(result) == 0 {
		return &desc.CalculateResponse{}
	}

	items := make([]*desc.CalculateResponse_Item, 0, len(result))

	for _, res := range result {
		item := &desc.CalculateResponse_Item{
			Var:   res.Var,
			Value: strconv.FormatInt(res.Value, 10),
		}

		items = append(items, item)
	}

	return &desc.CalculateResponse{
		Items: items,
	}
}

var exprType = map[string]model.Type{
	"calc":  model.Calc,
	"print": model.Print,
}

func convertTypeToModel(t string) model.Type {
	if v, ok := exprType[t]; ok {
		return v
	}

	return model.UnknownType
}

var exprOperation = map[string]model.Operation{
	"+": model.Addition,
	"-": model.Subtraction,
	"*": model.Multiplication,
	"/": model.Division,
}

func convertOperationToModel(o *string) *model.Operation {
	if o == nil {
		return nil
	}

	if v, ok := exprOperation[*o]; ok {
		return &v
	}

	return lo.ToPtr(model.UnknownOperation)
}
