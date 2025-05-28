package converter

import (
	"calculator/internal/pkg/model"
	desc "calculator/pkg/api"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func Test_ExpressionToModel(t *testing.T) {
	tests := []struct {
		name     string
		input    *desc.CalculateRequest
		expected []*model.Expression
	}{
		{
			name: "Valid input with literal operands",
			input: &desc.CalculateRequest{
				Data: []*desc.CalculateRequest_Data{
					{
						Type:       "calc",
						Op:         lo.ToPtr("+"),
						Var:        "x",
						LeftValue:  &desc.CalculateRequest_Data_LeftConst{LeftConst: 1},
						RightValue: &desc.CalculateRequest_Data_RightConst{RightConst: 2},
					},
				},
			},
			expected: []*model.Expression{
				{
					Type:  model.Calc,
					Op:    lo.ToPtr(model.Addition),
					Var:   "x",
					Left:  model.RawOperand{IsVar: false, IntVal: 1, VarName: ""},
					Right: model.RawOperand{IsVar: false, IntVal: 2, VarName: ""},
				},
			},
		},
		{
			name: "Valid input with variable operands",
			input: &desc.CalculateRequest{
				Data: []*desc.CalculateRequest_Data{
					{
						Type:       "calc",
						Op:         lo.ToPtr("*"),
						Var:        "y",
						LeftValue:  &desc.CalculateRequest_Data_LeftVar{LeftVar: "x"},
						RightValue: &desc.CalculateRequest_Data_RightVar{RightVar: "5"},
					},
				},
			},
			expected: []*model.Expression{
				{
					Type:  model.Calc,
					Op:    lo.ToPtr(model.Multiplication),
					Var:   "y",
					Left:  model.RawOperand{IsVar: true, IntVal: 0, VarName: "x"},
					Right: model.RawOperand{IsVar: true, IntVal: 0, VarName: "5"},
				},
			},
		},
		{
			name: "Print instruction without operands",
			input: &desc.CalculateRequest{
				Data: []*desc.CalculateRequest_Data{
					{Type: "print", Op: nil, Var: "z"},
				},
			},
			expected: []*model.Expression{
				{Type: model.Print, Op: nil, Var: "z", Left: model.RawOperand{}, Right: model.RawOperand{}},
			},
		},
		{
			name:     "Empty input",
			input:    &desc.CalculateRequest{Data: []*desc.CalculateRequest_Data{}},
			expected: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ExpressionToModel(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestResultToProto(t *testing.T) {
	tests := []struct {
		name     string
		input    []*model.Result
		expected *desc.CalculateResponse
	}{
		{
			name: "Valid result list",
			input: []*model.Result{
				{Var: "a", Value: 10},
				{Var: "b", Value: -5},
			},
			expected: &desc.CalculateResponse{
				Items: []*desc.CalculateResponse_Item{
					{Var: "a", Value: "10"},
					{Var: "b", Value: "-5"},
				},
			},
		},
		{
			name:     "Nil input",
			input:    nil,
			expected: &desc.CalculateResponse{},
		},
		{
			name:     "Empty slice",
			input:    []*model.Result{},
			expected: &desc.CalculateResponse{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ResultToProto(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
