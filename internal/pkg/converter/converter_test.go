package converter

import (
	"calculator/internal/pkg/model"
	desc "calculator/pkg/api"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ExpressionToModel(t *testing.T) {
	tests := []struct {
		name     string
		input    *desc.CalculateRequest
		expected []*model.Expression
	}{
		{
			name: "Valid input with calc and operation",
			input: &desc.CalculateRequest{
				Data: []*desc.CalculateRequest_Data{
					{
						Type:  "calc",
						Op:    lo.ToPtr("+"),
						Var:   "x",
						Left:  lo.ToPtr("1"),
						Right: lo.ToPtr("2"),
					},
				},
			},
			expected: []*model.Expression{
				{
					Type:  model.Calc,
					Op:    lo.ToPtr(model.Addition),
					Var:   "x",
					Left:  lo.ToPtr("1"),
					Right: lo.ToPtr("2"),
				},
			},
		},
		{
			name: "Unknown type and operation",
			input: &desc.CalculateRequest{
				Data: []*desc.CalculateRequest_Data{
					{
						Type:  "unknown",
						Op:    lo.ToPtr("%"),
						Var:   "y",
						Left:  lo.ToPtr("3"),
						Right: lo.ToPtr("4"),
					},
				},
			},
			expected: []*model.Expression{
				{
					Type:  model.UnknownType,
					Op:    lo.ToPtr(model.UnknownOperation),
					Var:   "y",
					Left:  lo.ToPtr("3"),
					Right: lo.ToPtr("4"),
				},
			},
		},
		{
			name: "Nil operation",
			input: &desc.CalculateRequest{
				Data: []*desc.CalculateRequest_Data{
					{
						Type:  "print",
						Op:    nil,
						Var:   "z",
						Left:  lo.ToPtr("0"),
						Right: lo.ToPtr("0"),
					},
				},
			},
			expected: []*model.Expression{
				{
					Type:  model.Print,
					Op:    nil,
					Var:   "z",
					Left:  lo.ToPtr("0"),
					Right: lo.ToPtr("0"),
				},
			},
		},
		{
			name: "Empty input",
			input: &desc.CalculateRequest{
				Data: []*desc.CalculateRequest_Data{},
			},
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
