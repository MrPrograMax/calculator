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
		var leftOp model.RawOperand
		switch lv := expr.LeftValue.(type) {
		case *desc.CalculateRequest_Data_LeftConst:
			leftOp = model.RawOperand{
				IsVar:  false,
				IntVal: lv.LeftConst,
			}
		case *desc.CalculateRequest_Data_LeftVar:
			leftOp = model.RawOperand{
				IsVar:   true,
				VarName: lv.LeftVar,
			}
		default:
			// для print-пакета left останется zero-value (IsVar=false, IntVal=0)
		}

		var rightOp model.RawOperand
		switch rv := expr.RightValue.(type) {
		case *desc.CalculateRequest_Data_RightConst:
			rightOp = model.RawOperand{
				IsVar:  false,
				IntVal: rv.RightConst,
			}
		case *desc.CalculateRequest_Data_RightVar:
			rightOp = model.RawOperand{
				IsVar:   true,
				VarName: rv.RightVar,
			}
		default:
			// для print-пакета right останется zero-value
		}

		expressions = append(expressions, &model.Expression{
			Type:  convertTypeToModel(expr.Type),
			Op:    convertOperationToModel(expr.Op),
			Var:   expr.Var,
			Left:  leftOp,
			Right: rightOp,
		})
	}

	return expressions
}

func ResultToProto(result []*model.Result) *desc.CalculateResponse {
	if len(result) == 0 {
		return &desc.CalculateResponse{}
	}
	items := make([]*desc.CalculateResponse_Item, 0, len(result))
	for _, res := range result {
		items = append(items, &desc.CalculateResponse_Item{
			Var:   res.Var,
			Value: strconv.FormatInt(res.Value, 10),
		})
	}
	return &desc.CalculateResponse{Items: items}
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
